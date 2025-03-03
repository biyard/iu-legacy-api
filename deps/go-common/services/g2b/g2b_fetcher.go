package g2b

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"biyard.co/common/base"
	types "biyard.co/common/base"
	"biyard.co/common/database/dynamodb"
)

// FIXME: This package has to be moved to `biyard.co/finder-api` package because this
// package is not a common package. It means that this package is dedicated to G2B fetcher.

type G2BFetcher struct {
	ctx context.Context
	db  *dynamodb.DynamoDB
	// announces  *dynamodb.Table[announce.Announce]
	// webhooks   *dynamodb.Table[announce.WebHook]
	// fetchindex *dynamodb.Table[announce.FetcherIndex]

	fstart string
	ftype  string
}

func NewG2BFetcher(ctx context.Context, cfg base.Config, fetchertype string, startdate string) *G2BFetcher {
	db := dynamodb.New(ctx, cfg.Database, cfg.AWS)
	// announces, err := dynamodb.NewTable[announce.Announce](ctx, db)
	// if err != nil {
	// 	panic(err)
	// }

	// indexes, err := dynamodb.NewTable[announce.FetcherIndex](ctx, db)
	// if err != nil {
	// 	panic(err)
	// }

	// webhooks, err := dynamodb.NewTable[announce.WebHook](ctx, db)
	// if err != nil {
	// 	panic(err)
	// }

	cli := &G2BFetcher{
		ctx: ctx,
		db:  db,
		// announces:  announces,
		// webhooks:   webhooks,
		// fetchindex: indexes,
		fstart: startdate,
		ftype:  fetchertype,
	}

	// _, err = cli.fetchindex.FindOne(ctx, bson.D{{Key: "key", Value: cli.IndexKey()}})
	// // logger.Debugf(ctx, "g2b fetcher err: %v doc: %+v", err, doc.Indexes())
	// if err != nil {
	// 	if err := cli.fetchindex.Insert(ctx, &announce.FetcherIndex{
	// 		Key:       cli.IndexKey(),
	// 		CreatedAt: time.Now().Unix(),
	// 		Type:      cli.ftype,
	// 		StartDate: cli.fstart,
	// 	}); err != nil {
	// 		panic(err)
	// 	}
	// }

	return cli
}

func (r *G2BFetcher) IndexKey() string {
	return fmt.Sprintf("%s#%s", r.ftype, r.fstart)
}

func (r *G2BFetcher) RunTask() *types.Error {
	// now = time.Now().Format("2006/01/02")
	// nowTime := stringToTime(now)

	// for {
	// 	dt, _ := mongoindex.FindOne(cont, bson.D{{Key: "type", Value: ftype}})
	// 	basetime := dt.StartDate
	// 	starttime := dt.StartDate
	// 	start := stringToTime(starttime)
	// 	if nowTime.Unix() <= start.Unix() {
	// 		break
	// 	}
	// 	end := start.AddDate(0, 0, 1)
	// 	basetime = end.Format("2006/01/02")
	// 	endtime := basetime

	// 	starturl := stringToUri(starttime)
	// 	endurl := stringToUri(endtime)

	// 	page := 1
	// 	for {
	// 		pagestring := strconv.Itoa(page)
	// 		url := `https://www.g2b.go.kr:8101/ep/tbid/tbidList.do?area=&bidNm=&bidSearchType=1&fromBidDt=` + starturl + `&fromOpenBidDt=&instNm=&maxPageViewNoByWshan=2&radOrgan=1&regYn=Y&searchDtType=1&searchType=1&taskClCds=&toBidDt=` + endurl + `&toOpenBidDt=&currentPageNo=` + pagestring
	// 		prevdata, _ := mongoannounces.Find(cont, bson.D{})
	// 		parsingData(url, nowTime.Unix())
	// 		nextdata, _ := mongoannounces.Find(cont, bson.D{})

	// 		if len(prevdata) == len(nextdata) {
	// 			break
	// 		}
	// 		page++
	// 		time.Sleep(100 * time.Millisecond)
	// 	}

	// 	_, err := mongoindex.UpdateOne(cont, bson.D{{Key: "type", Value: ftype}}, bson.D{{Key: "$set", Value: bson.D{{Key: "startdate", Value: endtime}}}})
	// 	if err != nil {
	// 		logger.Debugf(cont, "error happening")
	// 		return errs.ErrFailedToUpdateAddress
	// 	}

	// 	time.Sleep(1 * time.Second)
	// }

	// logger.Debugf(cont, "This function call1")

	// data, _ := mongoannounces.Find(cont, bson.D{{Key: "startdate", Value: nowTime.Unix()}})
	// webhook, _ := mongowebhook.Find(cont, bson.D{{}})

	// for ind := range webhook {
	// 	for i := range data {
	// 		endate := time.Unix(data[i].EndDate, 0).Format("2006/01/02")
	// 		if strings.Contains(data[i].Title, "블록체인") {
	// 			sendSlack(webhook[ind].Webhook, data[i].Title, data[i].Price, data[i].Link, endate)
	// 		} else if strings.Contains(data[i].Title, "플랫폼") && strings.Contains(data[i].Title, "구축") {
	// 			sendSlack(webhook[ind].Webhook, data[i].Title, data[i].Price, data[i].Link, endate)
	// 		} else if strings.Contains(data[i].Title, "소프트웨어") && strings.Contains(data[i].Title, "개발") {
	// 			sendSlack(webhook[ind].Webhook, data[i].Title, data[i].Price, data[i].Link, endate)
	// 		} else if strings.Contains(data[i].Title, "플랫폼") && strings.Contains(data[i].Title, "개발") {
	// 			sendSlack(webhook[ind].Webhook, data[i].Title, data[i].Price, data[i].Link, endate)
	// 		} else if strings.Contains(data[i].Title, "공고") {
	// 			sendSlack(webhook[ind].Webhook, data[i].Title, data[i].Price, data[i].Link, endate)
	// 		}
	// 	}
	// }

	return &types.Error{}
}

func sendSlack(webhook string, title string, price int, link string, enddate string) {
	slackPayload := Payload{
		Text: title + " " + strconv.Itoa(price) + "원(-1: 미정) (마감: " + enddate + ") " + `<` + link + `|Link>`,
	}
	slackPayload.SendSlack(webhook)
}

// func parsingData(baseUrl string, nowTime int64) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			debug.PrintStack()
// 			err := errors.New(fmt.Sprint(r))
// 			logger.Debugf(cont, "error: ", err)
// 		}
// 	}()

// 	url := baseUrl

// 	announceModels := make([]announce.Announce, 0)
// 	// res, _ := http.Get(url)
// 	acquireRequest := fasthttp.AcquireRequest()
// 	defer fasthttp.ReleaseRequest(acquireRequest)
// 	res := fasthttp.AcquireResponse()
// 	defer fasthttp.ReleaseResponse(res)

// 	acquireRequest.SetRequestURI(url)

// 	if err := fasthttp.Do(acquireRequest, res); err != nil {
// 		log.Panic(err)
// 	}

// 	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))

// 	links := doc.Find(".table_list_tbidTbl tbody tr")

// 	links.EachWithBreak(func(i int, s *goquery.Selection) bool {
// 		typ := s.Find("td").Text()
// 		totaldata, totalerror := iconv.ConvertString(string(typ), "euc-kr", "utf-8")

// 		if totalerror != nil {
// 			logger.Debugf(context.TODO(), baseUrl)
// 			return false
// 		}

// 		aTag := s.Find("a")
// 		href, _ := aTag.Attr("href")

// 		title := aTag.Text()
// 		out, outerror := iconv.ConvertString(string(title), "euc-kr", "utf-8")

// 		if outerror != nil {
// 			logger.Debugf(context.TODO(), baseUrl)
// 			return false
// 		}

// 		stdate := s.Find(".tc").Text()
// 		endate := s.Find(".tc .blue").Text()
// 		out2, out2error := iconv.ConvertString(string(stdate), "euc-kr", "utf-8")

// 		if out2error != nil || len(out2) == 0 {
// 			logger.Debugf(context.TODO(), baseUrl)
// 			return false
// 		}

// 		start := out2[:16]
// 		out3, out3error := iconv.ConvertString(string(endate), "euc-kr", "utf-8")

// 		if out3error != nil || len(out3) == 0 {
// 			logger.Debugf(context.TODO(), baseUrl)
// 			return false
// 		}

// 		end := out3[1 : len(out3)-1]

// 		startdate, enddate := parsingDate(start, end)

// 		re := regexp.MustCompile("[0-9a-zA-Z]+")
// 		iddata := re.FindAllString(out, -1)
// 		id := out[:len(iddata[0])+3]

// 		typedata := totaldata[0:6]
// 		classification := totaldata[6+len(id) : 12+len(id)]
// 		link := strings.Trim(href, " ")

// 		announcementagency, demandagency, manager, pri, bailiff := getInnerAnnounceData(link)

// 		announceModels = append(announceModels, announce.Announce{
// 			Key:                id + "(" + ftype + ")",
// 			Id:                 id,
// 			Type:               typedata,
// 			Classification:     classification,
// 			Link:               link,
// 			AnnouncementAgency: announcementagency,
// 			DemandAgency:       demandagency,
// 			Manager:            manager,
// 			Bailiff:            bailiff,
// 			Title:              out[len(id):],
// 			Price:              pri,
// 			StartDate:          startdate.Unix(),
// 			EndDate:            enddate.Unix(),
// 		})

// 		return true
// 	})

// 	logger.Debugf(cont, "dt: ", baseUrl)

// 	if len(announceModels) != 0 {
// 		err := mongoannounces.InsertMany(cont, announceModels)
// 		if err != nil {
// 			logger.Errorf(cont, "insert error: %v", err)
// 		}
// 	}
// }

// func getInnerAnnounceData(url string) (announcement string, demand string, manager string, price int, bailiff string) {
// 	ind := strings.LastIndex(url, "https")
// 	if ind == -1 {
// 		ind = 0
// 	}

// 	baseurl := url[ind:]

// 	acquireRequest := fasthttp.AcquireRequest()
// 	defer fasthttp.ReleaseRequest(acquireRequest)
// 	res := fasthttp.AcquireResponse()
// 	defer fasthttp.ReleaseResponse(res)

// 	acquireRequest.SetRequestURI(baseurl)

// 	if err := fasthttp.Do(acquireRequest, res); err != nil {
// 		log.Panic(err)
// 	}

// 	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
// 	links := doc.Find(".table_info tbody tr")

// 	in := 0

// 	announcementagency := ""
// 	demandagency := ""

// 	managerdt := ""
// 	bailiffdt := ""
// 	pri := -1

// 	doc.Find(".detail .table_info tbody tr").Each(func(i int, s *goquery.Selection) {
// 		dt := s.Find("td .tb_inner").Text()
// 		pricedt, _ := iconv.ConvertString(string(dt), "euc-kr", "utf-8")

// 		part := strings.Split(pricedt, "원")
// 		logger.Debugf(context.TODO(), "hi1:", pricedt)

// 		for _, v := range part {
// 			match, _ := regexp.MatchString("￦\\d{1,3}(,\\d{3})", v)

// 			logger.Debugf(context.TODO(), "hi2:", match)

// 			if match == true {
// 				d := v[3:]
// 				parsedt := strings.Replace(d, ",", "", -1)
// 				logger.Debugf(context.TODO(), "hello:", parsedt)
// 				pc, _ := strconv.Atoi(parsedt)
// 				if pri < pc {
// 					pri = pc
// 				}
// 			}
// 		}
// 	})

// 	if pri == -1 {
// 		doc.Find(".section .table_info").Each(func(i int, s *goquery.Selection) {
// 			dt, _ := s.Attr("summary")
// 			st, _ := iconv.ConvertString(string(dt), "euc-kr", "utf-8")
// 			if st == "예정가격 결정 및 입찰금액 정보" {
// 				s.Find(".tb_inner").Each(func(i int, s *goquery.Selection) {
// 					d := s.Text()
// 					priced, _ := iconv.ConvertString(string(d), "euc-kr", "utf-8")
// 					check := strings.Contains(priced, "원")

// 					if check {
// 						re := regexp.MustCompile("\\d{1,3}(,\\d{3})+원")
// 						part := re.FindAllString(priced, -1)
// 						for _, v := range part {
// 							match, _ := regexp.MatchString("\\d{1,3}(,\\d{3})+원", v)

// 							if match == true {
// 								d := v[:len(v)-3]
// 								parsedt := strings.Replace(d, ",", "", -1)
// 								pc, _ := strconv.Atoi(parsedt)
// 								if pri < pc {
// 									pri = pc
// 								}
// 							}
// 						}
// 					}
// 				})
// 			}
// 		})
// 	}

// 	links.Each(func(i int, s *goquery.Selection) {
// 		typ := s.Find("td .tb_inner .darkblue").Text()
// 		totaldata, _ := iconv.ConvertString(string(typ), "euc-kr", "utf-8")

// 		if totaldata != "" {
// 			s.Find("td .tb_inner").Children().Each(func(i int, s *goquery.Selection) {
// 				dt, _ := iconv.ConvertString(s.Text(), "euc-kr", "utf-8")

// 				if in == 0 && i == 0 {
// 					announcementagency = dt
// 				}
// 				if in == 0 && i == 1 {
// 					demandagency = dt
// 				}
// 				if in == 1 && i == 0 {
// 					managerdt = dt
// 				}
// 				if in == 1 && i == 1 {
// 					bailiffdt = dt
// 				}
// 			})

// 			in++
// 		}
// 	})

// 	return announcementagency, demandagency, managerdt, pri, bailiffdt
// }

func parsingDate(startdate string, enddate string) (std time.Time, en time.Time) {
	var styear int = 0
	var stmonth int = 0
	var stday int = 0

	if startdate != "-" {
		startyear, _ := strconv.Atoi(startdate[:4])
		startmonth, _ := strconv.Atoi(startdate[5:7])
		startday, _ := strconv.Atoi(startdate[8:10])

		styear = startyear
		stmonth = startmonth
		stday = startday
	}

	var enyear int = 0
	var enmonth int = 0
	var enday int = 0

	if enddate != "-" {
		endyear, _ := strconv.Atoi(enddate[:4])
		endmonth, _ := strconv.Atoi(enddate[5:7])
		endday, _ := strconv.Atoi(enddate[8:10])

		enyear = endyear
		enmonth = endmonth
		enday = endday
	}

	start := time.Date(styear, time.Month(stmonth), stday, 0, 0, 0, 0, time.UTC)
	end := time.Date(enyear, time.Month(enmonth), enday, 0, 0, 0, 0, time.UTC)
	return start, end
}

func stringToTime(t string) time.Time {
	delimiter := "/"
	part := strings.Split(t, delimiter)
	year, _ := strconv.Atoi(part[0])
	month, _ := strconv.Atoi(part[1])
	day, _ := strconv.Atoi(part[2])
	ti := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return ti
}

func stringToUri(t string) string {
	delimiter := "/"
	part := strings.Split(t, delimiter)
	return part[0] + "%2F" + part[1] + "%2F" + part[2]
}
