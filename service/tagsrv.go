// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"sync"

	"github.com/b3log/pipe/model"
)

// Tag service.
var Tag = &tagService{
	mutex: &sync.Mutex{},
}

type tagService struct {
	mutex *sync.Mutex
}

func (srv *tagService) GetTags(size int, blogID uint64) (ret []*model.Tag) {
	if err := db.Where("`blog_id` = ?", blogID).Order("`article_count` DESC, `id` DESC").Limit(size).Find(&ret).Error; nil != err {
		logger.Errorf("get tags failed: " + err.Error())
	}

	return
}

func (srv *tagService) GetTagByTitle(title string, blogID uint64) *model.Tag {
	ret := &model.Tag{}
	if err := db.Where("`title` = ? AND `blog_id` = ?", title, blogID).First(ret).Error; nil != err {
		return nil
	}

	return ret
}
