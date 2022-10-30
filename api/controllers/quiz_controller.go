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

  var countries []models.Country
  var flag_colors []models.FlagColor
  var results []models.Result

  // 初めに全ての国を候補とする
  options := mapset.NewSet[string]()
  database.DB.Find(&countries)
  for _, v := range countries {
    options.Add(v.ID)
  }

  // 地域が該当する国を絞る
  if len(areas) != 0 {
    database.DB.Where("area_id in (?)", areas).Find(&countries)
    selected := mapset.NewSet[string]()
    for _, v := range countries {
      selected.Add(v.ID)
    }
    options = options.Intersect(selected)
  }

  // 色が完全一致する国を絞る
  if len(colors) != 0 {
    database.DB.Select("country_id").Where("color_id in (?)", colors).
      Group("country_id").Having("count(*) = ?", len(colors)).Find(&flag_colors)
    selected := mapset.NewSet[string]()
    for _, v := range flag_colors {
      selected.Add(v.CountryID)
    }
    options = options.Intersect(selected)
  }

  // 該当ユーザーがブックマークした国を絞る
  if  (err == nil) && (bookmark == 1) {
    database.DB.Where("user_id=?", user.ID).Where("bookmark=?", bookmark).Find(&results)
    selected := mapset.NewSet[string]()
    for _, v := range results {
      selected.Add(v.CountryID)
    }
    options = options.Intersect(selected)
  }

  log.Println(options)

  // 絞り込んだ国からクイズを選択
  var quizzes []models.Quiz

  c.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}