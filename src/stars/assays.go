package main

import (
	. "domain"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"
	. "repository/regulators"
	"time"
)

func AssaysPage(rd render.Render) {
	rd.HTML(http.StatusOK, "pages/assays", nil)
}

type VisitsModel struct {
	Login  string        `json:"login"`
	Visits []*VisitModel `json:"visits"`
}

type VisitModel struct {
	Assay   string `json:"assay"`
	Period  string `json:"period"`
	Active  bool   `json:"active"`
	Visited string `json:"visited"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func AssaysList(rd render.Render, regulators IRegulatorsRepository, params martini.Params) {
	login := params["login"]
	if len(login) == 0 {
		rd.Error(http.StatusBadRequest)
		return
	}
	visits, err := regulators.AllEmployeeAssaysLastVisits(login)
	if err != nil {
		rd.Error(http.StatusBadRequest)
		return
	}

	list := make([]*VisitModel, 0, len(visits))
	for _, v := range visits {
		var visited string
		if !v.Visited().IsZero() {
			visited = v.Visited().Format(DateShortFormat)
		}
		list = append(list, &VisitModel{
			Assay:   v.AssayId(),
			Period:  fmt.Sprintf("%v", v.AssayPeriod()),
			Active:  v.Active(),
			Visited: visited,
			Status:  v.Status(),
			Message: v.Message(),
		})
	}
	model := VisitsModel{
		Login:  login,
		Visits: list,
	}
	rd.JSON(http.StatusOK, model)
}

//
func AssaysUpdate(rd render.Render, regulators IRegulatorsRepository, m VisitsModel) {
	visits, err := regulators.AllEmployeeAssaysLastVisits(m.Login)
	if GotError(err, rd) {
		return
	}
	for i, v := range visits {
		vm, found, err := From(m.Visits).FirstBy(func(b T) (bool, error) {
			return b.(*VisitModel).Assay == v.AssayId(), nil
		})
		if GotError(err, rd) {
			return
		}
		if found {
			model := vm.(*VisitModel)
			var visited time.Time
			if len(model.Visited) > 0 {
				visited, err = time.Parse(DateShortFormat, model.Visited)
				if GotError(err, rd) {
					return
				}
			}
			rec := visits[i]
			rec.Active(model.Active)
			rec.Visited(visited)
		}
	}
	fmt.Println(visits)
	err = regulators.UpdateEmployeeAssaysLastVisits(m.Login, visits)
	if GotError(err, rd) {
		return
	}
	rd.JSON(http.StatusOK, OperationResponse{
		Status: StatusSuccess,
	})
}
