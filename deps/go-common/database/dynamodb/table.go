package dynamodb

import (
	"context"
	"encoding"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.mongodb.org/mongo-driver/bson"
)

type ITable interface {
	UpdateOne(ctx context.Context, filter any, doc any) error
}

type TableDescriber interface {
	Table() string
}

type Table[T TableDescriber] struct {
	tableName *string
	svc       *dynamodb.DynamoDB
}

func NewTable[T TableDescriber](ctx context.Context, db *DynamoDB) (*Table[T], error) {
	var doc T
	n := fmt.Sprintf("%s-%s", db.prefix, doc.Table())

	return &Table[T]{
		tableName: &n,
		svc:       db.svc,
	}, nil
}

func (r *Table[T]) FindAll(ctx context.Context) ([]T, error) {
	out, _ := r.svc.Scan(&dynamodb.ScanInput{
		TableName: r.tableName,
		Limit:     aws.Int64(1000),
	})

	var doc []T
	if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &doc); err != nil {
		return nil, err
	}

	return doc, nil
}

func (r *Table[T]) FindOne(ctx context.Context, filter any) (*T, *base.Error) {
	var doc T

	input := &dynamodb.GetItemInput{
		TableName: r.tableName,
		Key:       compositeKey(filter),
	}
	result, e := r.svc.GetItem(input)
	if e != nil {
		return nil, base.ErrDatabaseUnknown.WithMessage(e.Error())
	}

	if len(result.Item) == 0 {
		return nil, base.ErrDataNotFound
	}

	logger.Debugf(ctx, "result: %+v, len: %v", result, len(result.Item))

	if e = dynamodbattribute.UnmarshalMap(result.Item, &doc); e != nil {
		return nil, base.ErrDataParsing.WithMessage(e.Error())
	}

	return &doc, nil
}

func (r *Table[T]) Find(ctx context.Context, filter any) ([]T, error) {
	// TODO: scan-based find but not recommended
	docs := make([]T, 0)

	return docs, nil
}

func (r *Table[T]) DeleteOne(ctx context.Context, filter any) error {
	key := compositeKey(filter)

	dt := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: r.tableName,
	}

	_, err := r.svc.DeleteItem(dt)
	return err
}

func (r *Table[T]) UpdateOne(ctx context.Context, filter any, doc any) error {
	key := compositeKey(filter)
	if v, ok := doc.(bson.D); ok {
		doc = v[0].Value
	}

	var result map[string]*dynamodb.AttributeValue
	var cond string
	var err error

	if v, ok := doc.(bson.D); ok {
		result, cond = compositeValuesWithBson(v)
	} else {
		result, err = dynamodbattribute.MarshalMap(doc)
		if err != nil {
			return err
		}

	}
	logger.Debugf(ctx, "result: %+v \n %+v", key, result)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: result,
		Key:                       key,
		UpdateExpression:          aws.String(fmt.Sprintf("SET %s", cond)),
		TableName:                 r.tableName,
	}

	_, err = r.svc.UpdateItemWithContext(ctx, input)

	return err
}

type Polisher interface {
	Polish()
}

func (r *Table[T]) Insert(ctx context.Context, doc *T) error {
	if v, ok := reflect.ValueOf(&doc).Interface().(Polisher); ok {
		v.Polish()
	}

	result, err := dynamodbattribute.MarshalMap(doc)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      result,
		TableName: r.tableName,
	}

	_, err = r.svc.PutItem(input)

	return err
}

func (r *Table[T]) InsertMany(ctx context.Context, docs []*T) error {
	for _, doc := range docs {
		m := reflect.ValueOf(doc).MethodByName("Polish")
		if m.IsValid() {
			m.Call([]reflect.Value{})
		}
	}

	req := map[string][]*dynamodb.WriteRequest{}
	for _, doc := range docs {
		item, err := dynamodbattribute.MarshalMap(doc)
		if err != nil {
			return err
		}

		req[*r.tableName] = append(req[*r.tableName], &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: item,
			},
		})
	}

	_, err := r.svc.BatchWriteItemWithContext(
		ctx,
		&dynamodb.BatchWriteItemInput{
			RequestItems: req,
		},
	)

	return err
}

func parseElem(val reflect.Value) *dynamodb.AttributeValue {
	var attrVal *dynamodb.AttributeValue

	switch v := val.Interface().(type) {
	case encoding.TextMarshaler:
		if text, err := v.MarshalText(); err == nil {
			return &dynamodb.AttributeValue{S: aws.String(string(text))}
		}
	case dynamodbattribute.Marshaler:
		if marshal, err := dynamodbattribute.Marshal(val.Interface()); err == nil {
			return marshal
		}
	case dynamodb.AttributeValue:
		return &v
	}

	switch val.Kind() {
	case reflect.Bool:
		attrVal = &dynamodb.AttributeValue{
			BOOL: aws.Bool(val.Bool()),
		}
	case reflect.String:
		attrVal = &dynamodb.AttributeValue{
			S: aws.String(val.String()),
		}
	case reflect.Uint8, reflect.Uint16:
		attrVal = &dynamodb.AttributeValue{
			N: aws.String(strconv.Itoa(int(val.Uint()))),
		}
	case reflect.Uint64:
		attrVal = &dynamodb.AttributeValue{
			N: aws.String(strconv.FormatUint(val.Uint(), 10)),
		}
	case reflect.Int, reflect.Int32:
		attrVal = &dynamodb.AttributeValue{
			N: aws.String(strconv.Itoa(int(val.Int()))),
		}
	case reflect.Int64:
		attrVal = &dynamodb.AttributeValue{
			N: aws.String(strconv.FormatInt(val.Int(), 10)),
		}
	case reflect.Float64:
		attrVal = &dynamodb.AttributeValue{
			N: aws.String(strconv.FormatFloat(math.Abs(val.Float()), 'f', -1, 64)),
		}
	case reflect.Array, reflect.Slice:
		l := val.Len()
		lVal := []*dynamodb.AttributeValue{}

		for i := 0; i < l; i++ {
			v := val.Index(i)
			vi := &dynamodb.AttributeValue{}

			if reflect.Indirect(v).Kind() == reflect.Struct {
				vi.M = parseObjectDocument(v.Interface())
			} else {
				vi = parseElem(v)
			}
			lVal = append(lVal, vi)
		}
		attrVal = &dynamodb.AttributeValue{
			L: lVal,
		}
	case reflect.Struct:
		attrVal = &dynamodb.AttributeValue{
			M: parseObjectDocument(val.Interface()),
		}
	case reflect.Interface:
		if val.IsNil() {
			return nil
		}
		return parseElem(val.Elem())
	case reflect.Map:
		m, ok := val.Interface().(map[string]interface{})
		if !ok {
			panic(base.ErrUnsupportedType)
		}
		_, v, _ := prepareInputFromMap(m)
		attrVal = &dynamodb.AttributeValue{
			M: v,
		}
	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}
		return parseElem(reflect.Indirect(val))
	default:
		panic(base.ErrUnsupportedType)
	}

	return attrVal
}

func prepareInputFromMap(doc map[string]interface{}) ([]string, map[string]*dynamodb.AttributeValue, map[string]*string) {
	updateExpr := []string{}
	exprAttrName := map[string]*string{}
	exprAttrVal := map[string]*dynamodb.AttributeValue{}

	for k, v := range doc {
		if v == nil {
			continue
		}
		kn := k
		replacer := strings.NewReplacer(".", "1", "<", "2", ">", "3", "#", "4")

		k = replacer.Replace(k)
		expr := fmt.Sprintf("#%s = :%s", k, k)
		attrVal := parseElem(reflect.ValueOf(v))
		if attrVal == nil {
			continue
		}

		exprAttrVal[fmt.Sprintf(":%s", k)] = attrVal
		l := kn
		exprAttrName[fmt.Sprintf("#%s", k)] = &l
		if !strings.HasPrefix(k, "v_") {
			updateExpr = append(updateExpr, expr)
		}
	}

	return updateExpr, exprAttrVal, exprAttrName
}

func parseObjectDocument(doc interface{}) map[string]*dynamodb.AttributeValue {
	exprAttrVal := map[string]*dynamodb.AttributeValue{}
	val := reflect.Indirect(reflect.ValueOf(doc))
	fieldNum := val.NumField()

	for i := 0; i < fieldNum; i++ {
		field := val.Field(i)
		ft := val.Type().Field(i)
		if ft.Name[0]&0x20 != 0 {
			continue
		}
		if ft.Anonymous {
			for k, v := range parseObjectDocument(field.Interface()) {
				if v != nil {
					exprAttrVal[fmt.Sprintf("%s", k)] = v
				}
			}

			continue
		}

		tag := ft.Tag.Get("json")
		required := ft.Tag.Get("required")
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = ft.Name
		}
		k := tag

		if required == "" && field.IsZero() {
			continue
		}
		v := parseElem(field)
		if v != nil {
			exprAttrVal[fmt.Sprintf("%s", k)] = v
		}
	}

	return exprAttrVal
}

func compositeKeyWithBson(filter bson.D) map[string]*dynamodb.AttributeValue {
	key := map[string]*dynamodb.AttributeValue{}

	for _, v := range filter {
		tag := strings.SplitN(v.Key, ".", 1)[0]
		key[tag] = parseElem(reflect.ValueOf(v.Value))
	}

	return key
}

func compositeValuesWithBson(filter bson.D) (map[string]*dynamodb.AttributeValue, string) {
	key := map[string]*dynamodb.AttributeValue{}
	conditions := []string{}

	for _, v := range filter {
		tag := strings.SplitN(v.Key, ".", 1)[0]
		key[":"+tag] = parseElem(reflect.ValueOf(v.Value))
		conditions = append(conditions, fmt.Sprintf("%s = :%s", tag, tag))
	}

	return key, strings.Join(conditions, ", ")
}

func compositeKey(doc any) map[string]*dynamodb.AttributeValue {
	if v, ok := doc.(bson.D); ok {
		return compositeKeyWithBson(v)
	}

	t := reflect.TypeOf(doc)
	v := reflect.ValueOf(doc)

	if t.Kind() != reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	key := map[string]*dynamodb.AttributeValue{}

	for i := 0; i < t.NumField(); i++ {
		val := v.Field(i)
		f := t.Field(i)
		if val == reflect.Zero(f.Type) {
			continue
		}
		if _, ok := f.Tag.Lookup("key"); ok {
			tag := strings.SplitN(f.Tag.Get("json"), ",", 1)[0]
			key[tag] = parseElem(val)
		}
	}

	return key
}
