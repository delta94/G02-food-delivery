package ginnote

import (
	"fooddlv/appctx"
	"fooddlv/common"
	"fooddlv/module/note/notebusiness"
	"fooddlv/module/note/notemodel"
	"fooddlv/module/note/notestorage"
	"github.com/gin-gonic/gin"
)

func CreateNote(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.NoteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := notestorage.NewSQLStore(db)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		bizNote := notebusiness.NewCreateNoteBiz(store, requester)

		data.UserId = requester.GetUserId()

		err := bizNote.CreateNewNote(c.Request.Context(), &data)

		//note, err := notebusiness.NewGetNoteBiz(store).GetNote(c.Request.Context(), data.Id)

		data.Mask(true)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId))
	}
}
