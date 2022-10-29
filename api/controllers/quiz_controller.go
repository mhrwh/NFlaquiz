package controllers

import (
  "log"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/raylicola/NFlaquiz/database"
  "github.com/raylicola/NFlaquiz/models"
  "github.com/raylicola/NFlaquiz/utils"
  mapset "github.com/deckarep/golang-set/v2"
)

// 検索条件に当てはまるクイズを10問選ぶ
// ヒット数が10個以下なら全て選択する
// 受信：
//   colors(array): 選択された色
//    areas(array): 選択された地域
//   bookmark(int): ブックマークで絞り込むか否か
// 戻り値：
//    10問以下のクイズセット
func SelectQuiz(c *gin.Context) {

  // 絞り込み条件を取得
  var req models.QuizFilter
  if err := c.Bind(&req); err != nil {
    c.JSON(http.StatusBadRequest, err)
    return
  }

  colors := req.Colors
  areas := req.Areas
  bookmark, _ := strconv.Atoi(req.Bookmark)
  user, err := utils.AuthUser(c)

  // 地域が該当する国
  if len(areas) == 0 {
    areas = append(areas, "Africa", "Asia", "Europe", "LatinAmericaandCaribbean",
                          "MiddleEast", "NorthAmerica", "Oceania")
  }
  var countries []models.Country
  database.DB.Where("area_id in (?)", areas).Find(&countries)
  options1 := mapset.NewSet[string]()
  for _, v := range countries {
    options1.Add(v.ID)
  }

  // 色が完全一致する国
  var flag_colors []models.FlagColor
  database.DB.Select("country_id").Where("color_id in (?)", colors).
    Group("country_id").Having("count(*) >= ?", len(colors)).Find(&flag_colors)
  options2 := mapset.NewSet[string]()
  for _, v := range flag_colors {
    options2.Add(v.CountryID)
  }

  // 該当ユーザーがブックマークした国
  var results []models.Result
  options3 := mapset.NewSet[string]()
  if  (err == nil) && (bookmark == 1) {
    database.DB.Where("user_id=?", user.ID).Where("bookmark=?", bookmark).Find(&results)
    for _, v := range results {
      options3.Add(v.CountryID)
    }
  } else {
    database.DB.Find(&countries)
    for _, v := range countries {
      options3.Add(v.ID)
    }
  }

  // 全ての絞り込み条件に一致した国
  options :=options1.Intersect(options2).Intersect(options3)
  log.Println(options)

  // クイズを選択
  var quizzes []models.Quiz

  c.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}