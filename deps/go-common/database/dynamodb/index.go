package dynamodb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndexDescriber interface {
	TableDescriber
	Indexes() []mongo.IndexModel
}

type Index[T IndexDescriber] struct {
	tableName *string
	indexName *string
	svc       *dynamodb.DynamoDB
	key       string
}

func NewIndex[T IndexDescriber](ctx context.Context, db *DynamoDB, indexName string) (*Index[T], error) {
	var doc T

	n := fmt.Sprintf("%s-%s", db.prefix, doc.Table())

	key, err := createIndex[T](ctx, db, n, indexName)
	if err != nil {
		return nil, err
	}

	return &Index[T]{
		tableName: &n,
		svc:       db.svc,
		indexName: &indexName,
		key:       key,
	}, nil
}

type QueryRequest struct {
	Key            string
	Type           int
	From           int64
	To             int64
	Bookmark       string
	Ascending      bool
	DisableSortKey bool
	IsNumber       bool
}

type QueryResult[T IndexDescriber] struct {
	Items    []T    `json:"items"`
	Bookmark string `json:"bookmark"`
}

func (r *Index[T]) FindAll(ctx context.Context, req QueryRequest) (*QueryResult[T], error) {
	ret := &QueryResult[T]{}

	ret, err := r.Find(ctx, req)
	if err != nil {
		return nil, err
	}

	req.Bookmark = ret.Bookmark

	for ret.Bookmark != "" {
		res, err := r.Find(ctx, req)
		if err != nil {
			return nil, err
		}

		logger.Debugf(ctx, "FindAll iteration result: %+v", len(res.Items))

		ret.Items = append(ret.Items, res.Items...)
		req.Bookmark = res.Bookmark
	}
	ret.Bookmark = ""

	return ret, nil
}

// bson
func (r *Index[T]) Find(ctx context.Context, req QueryRequest) (*QueryResult[T], error) {
	attrNames := map[string]*string{
		"#key": aws.String(r.key),
	}
	attrValues := map[string]*dynamodb.AttributeValue{
		":key": {
			S: aws.String(req.Key),
		},
	}

	if req.IsNumber {
		attrValues = map[string]*dynamodb.AttributeValue{
			":key": {
				N: aws.String(req.Key),
			},
		}
	}

	keyCondition := aws.String("#key = :key")

	if !req.DisableSortKey {

		attrNames["#createdAt"] = aws.String("createdAt")

		attrValues[":from"] = &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", req.From)),
		}

		if req.To == 0 {
			req.To = time.Now().Unix()
		}

		attrValues[":to"] = &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", req.To)),
		}

		keyCondition = aws.String("#key = :key and #createdAt BETWEEN :from AND :to")
	}

	input := &dynamodb.QueryInput{
		TableName:                 r.tableName,
		IndexName:                 r.indexName,
		Limit:                     aws.Int64(1000),
		ScanIndexForward:          &req.Ascending,
		ExpressionAttributeNames:  attrNames,
		ExpressionAttributeValues: attrValues,
		KeyConditionExpression:    keyCondition,
	}

	bookmark, err := r.DecodeBookmark(req.Bookmark)
	if err != nil {
		return nil, err
	}

	input.ExclusiveStartKey = bookmark

	out, err := r.svc.Query(input)
	logger.DebugH(ctx, func() {
		logger.Debugf(ctx, "dynamodb query: %+v", input)
		logger.Debugf(ctx, "dynamodb result: %+v", out)
	})

	if out == nil || err != nil {
		return nil, err
	}
	var doc []T
	if err := dynamodbattribute.UnmarshalListOfMaps(out.Items, &doc); err != nil {
		return nil, err
	}

	return &QueryResult[T]{
		Items:    doc,
		Bookmark: r.Bookmark(out.LastEvaluatedKey),
	}, nil
}

func (r *Index[T]) DecodeBookmark(b string) (map[string]*dynamodb.AttributeValue, error) {
	if b == "" {
		return nil, nil
	}

	bookmark, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return nil, err
	}

	var m map[string]*dynamodb.AttributeValue
	if err := json.Unmarshal(bookmark, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Index[T]) Bookmark(b map[string]*dynamodb.AttributeValue) string {
	if b == nil || len(b) == 0 {
		return ""
	}

	d, err := json.Marshal(&b)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(d)
}

// createIndex creates index if not exists. It decide the type of an attribute
// by Value of bson.E. If Value is `1`, it will be "S", and `-1` means "N".
func createIndex[T IndexDescriber](ctx context.Context, db *DynamoDB, n string, index string) (string, error) {
	var doc T

	key := ""
	keys := []*dynamodb.KeySchemaElement{}
	indexes := make([]*dynamodb.GlobalSecondaryIndexUpdate, 0)
	defs := make([]*dynamodb.AttributeDefinition, 0)

	for _, v := range doc.Indexes() {
		logger.Debugf(ctx, "option name: %+v %+v", *v.Options, index)
		if *v.Options.Name != index {
			continue
		}

		data, _ := v.Keys.(bson.D)
		key = data[0].Key
		// disableSortKey := data[0].Value != 1
		keys = []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(key),
				KeyType:       aws.String("HASH"),
			},
		}

		if data[0].Value == 6 {
			defs = []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String(key),
					AttributeType: aws.String("N"),
				},
			}
		} else {
			defs = []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String(key),
					AttributeType: aws.String("S"),
				},
			}
		}

		if data[0].Value == 1 || data[0].Value == 6 {
			keys = append(keys, &dynamodb.KeySchemaElement{
				AttributeName: aws.String("createdAt"),
				KeyType:       aws.String("RANGE"),
			})

			defs = append(defs, &dynamodb.AttributeDefinition{
				AttributeName: aws.String("createdAt"),
				AttributeType: aws.String("N"),
			})
		} else if data[0].Value == 2 {
			keys = append(keys, &dynamodb.KeySchemaElement{
				AttributeName: aws.String("dailyCount"),
				KeyType:       aws.String("RANGE"),
			})

			defs = append(defs, &dynamodb.AttributeDefinition{
				AttributeName: aws.String("dailyCount"),
				AttributeType: aws.String("N"),
			})
		} else if data[0].Value == 3 {
			keys = append(keys, &dynamodb.KeySchemaElement{
				AttributeName: aws.String("lv"),
				KeyType:       aws.String("RANGE"),
			})

			defs = append(defs, &dynamodb.AttributeDefinition{
				AttributeName: aws.String("lv"),
				AttributeType: aws.String("N"),
			})
		} else if data[0].Value == 4 {
			keys = append(keys, &dynamodb.KeySchemaElement{
				AttributeName: aws.String("exp"),
				KeyType:       aws.String("RANGE"),
			})

			defs = append(defs, &dynamodb.AttributeDefinition{
				AttributeName: aws.String("exp"),
				AttributeType: aws.String("N"),
			})
		} else if data[0].Value == 5 {
			keys = append(keys, &dynamodb.KeySchemaElement{
				AttributeName: aws.String("votingCount"),
				KeyType:       aws.String("RANGE"),
			})

			defs = append(defs, &dynamodb.AttributeDefinition{
				AttributeName: aws.String("votingCount"),
				AttributeType: aws.String("N"),
			})
		}

		logger.Debugf(ctx, "create index: %v", *v.Options.Name)
		indexes = append(indexes, &dynamodb.GlobalSecondaryIndexUpdate{
			Create: &dynamodb.CreateGlobalSecondaryIndexAction{
				IndexName: v.Options.Name,
				KeySchema: keys,
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
			},
		})

		break
	}

	d, err := db.svc.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: &n,
	})
	if err != nil {
		return key, base.ErrDatabaseUnknown.WithMessage(err.Error())
	}

	for _, v := range d.Table.GlobalSecondaryIndexes {
		if *v.IndexName == index {
			return key, nil
		}
	}

	if len(indexes) == 0 {
		return key, base.ErrDatabaseUnknown.WithMessage("index not found")
	}

	logger.Debugf(ctx, "update table: %+v", indexes)

	if _, err := db.svc.UpdateTable(&dynamodb.UpdateTableInput{
		TableName:                   &n,
		AttributeDefinitions:        defs,
		GlobalSecondaryIndexUpdates: indexes,
	}); err != nil {
		logger.Errorf(ctx, "Failed to create indexes: %+v", err)
		return key, base.ErrDatabaseUnknown.WithMessage(err.Error())
	}

	return key, nil
}
