package s2

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFromMap(t *testing.T) {
	Convey("Given invalid args", t, func() {
		Convey("When called", func() {
			err := FromMap(make(map[string]interface{}, 1), nil)
			Convey("Then should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given invalid args", t, func() {
		Convey("When called", func() {
			err := FromMap(nil, "")
			Convey("Then should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given valid struct", t, func() {

		type Fuga struct {
			City   string `json:"city"`
			Number int64
		}

		type Hoge struct {
			ID   string `json:"id"`
			Name string
			Age  int64
			Fuga Fuga
		}

		id := "id001"
		name := "name001"
		age := int64(88)
		city := "kasuya"
		number := int64(12345)

		s := Hoge{
			ID:   id,
			Name: name,
			Age:  age,
			Fuga: Fuga{
				City:   city,
				Number: number,
			},
		}

		m, _ := ToMap("json", s)

		Convey("When called", func() {
			hoge := &Hoge{}
			err := FromMap(m, hoge)

			Convey("Then should return map", func() {
				So(err, ShouldBeNil)
				So(hoge, ShouldNotBeNil)
				So(hoge.ID, ShouldEqual, id)
				So(hoge.Name, ShouldEqual, name)
				So(hoge.Age, ShouldEqual, age)

				So(hoge.Fuga.City, ShouldEqual, city)
				So(hoge.Fuga.Number, ShouldEqual, number)
			})
		})
	})
}

func TestToMap(t *testing.T) {
	Convey("Given invalid args", t, func() {
		Convey("When called", func() {
			m, err := ToMap("json", nil)
			Convey("Then should return error", func() {
				So(m, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})

	type Fuga struct {
		City   string `json:"city"`
		Number int64  `bson:"number,omitempty"`
	}

	type Hoge struct {
		ID   string `json:"id"`
		Name string `bson:"name"`
		Age  int64  `json:",string"`
		Fuga Fuga
	}

	id := "id001"
	name := "name001"
	age := int64(88)
	city := "kasuya"
	number := int64(12345)

	s := Hoge{
		ID:   id,
		Name: name,
		Age:  age,
		Fuga: Fuga{
			City:   city,
			Number: number,
		},
	}

	Convey("Given valid struct, no tagName", t, func() {
		Convey("When called", func() {
			m, err := ToMap("", s)

			Convey("Then should return map", func() {
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)
				So(len(m), ShouldEqual, 4)
				So(m["ID"], ShouldEqual, id)
				So(m["Name"], ShouldEqual, name)
				So(m["Age"], ShouldEqual, age)

				fuga, ok := m["Fuga"].(map[string]interface{})
				So(ok, ShouldBeTrue)
				So(fuga["City"], ShouldEqual, city)
				So(fuga["Number"], ShouldEqual, number)
			})
		})
	})

	Convey("Given valid struct, tagName bson", t, func() {
		Convey("When called", func() {
			m, err := ToMap("bson", s)

			Convey("Then should return map", func() {
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)
				So(len(m), ShouldEqual, 4)
				So(m["ID"], ShouldEqual, id)
				So(m["name"], ShouldEqual, name)
				So(m["Age"], ShouldEqual, age)

				fuga, ok := m["Fuga"].(map[string]interface{})
				So(ok, ShouldBeTrue)
				So(fuga["City"], ShouldEqual, city)
				So(fuga["number"], ShouldEqual, number)
			})
		})
	})

	Convey("Given valid struct, tagName json", t, func() {
		Convey("When called", func() {
			m, err := ToMap("json", s)

			Convey("Then should return map", func() {
				So(err, ShouldBeNil)
				So(m, ShouldNotBeNil)
				So(len(m), ShouldEqual, 4)
				So(m["id"], ShouldEqual, id)
				So(m["Name"], ShouldEqual, name)
				So(m["Age"], ShouldEqual, age)

				fuga, ok := m["Fuga"].(map[string]interface{})
				So(ok, ShouldBeTrue)
				So(fuga["city"], ShouldEqual, city)
				So(fuga["Number"], ShouldEqual, number)
			})
		})
	})
}
