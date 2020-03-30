package route

import (
	"bytes"
	"testing"

	"github.com/zhangchong5566/manba/pkg/pb/metapb"
)

func TestAddWithMethod(t *testing.T) {
	r := NewRoute()
	err := r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "GET",
	})
	if err != nil {
		t.Errorf("add error")
	}
	if len(r.root.children) != 1 {
		t.Errorf("expect 1 children but %d", len(r.root.children))
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "PUT",
	})
	if err != nil {
		t.Errorf("add error")
	}
	if len(r.root.children) != 1 {
		t.Errorf("expect 1 children but %d", len(r.root.children))
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "DELETE",
	})
	if err != nil {
		t.Errorf("add error")
	}
	if len(r.root.children) != 1 {
		t.Errorf("expect 1 children but %d", len(r.root.children))
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "POST",
	})
	if err != nil {
		t.Errorf("add error")
	}
	if len(r.root.children) != 1 {
		t.Errorf("expect 1 children but %d", len(r.root.children))
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "*",
	})
	if err == nil {
		t.Errorf("add error, expect error")
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/v1/users",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error")
	}

	err = r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/v1/users",
		Method:     "GET",
	})
	if err == nil {
		t.Errorf("add error, expect error")
	}
}

func TestMatchesSame(t *testing.T) {
	r := NewRoute()
	err := r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/(string):a/(string):b/(string):c",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error")
	}

	params := make(map[string]string)
	id, _ := r.Find([]byte("/a/b/c"), "GET", func(name, value []byte) {
		params[string(name)] = string(value)
	})
	if id != 1 {
		t.Errorf("expect matched 1, but %d", id)
	}

	if params["a"] != "a" {
		t.Errorf("expect param value a, but %s", params["a"])
	}
	if params["b"] != "b" {
		t.Errorf("expect param value b, but %s", params["b"])
	}
	if params["c"] != "c" {
		t.Errorf("expect param value c, but %s", params["c"])
	}
}

func TestAdd(t *testing.T) {
	r := NewRoute()
	err := r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/users",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error")
	}
	if len(r.root.children) != 1 {
		t.Errorf("expect 1 children but %d", len(r.root.children))
	}

	err = r.Add(&metapb.API{
		ID:         2,
		URLPattern: "/accounts",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error with 2 const")
	}
	if len(r.root.children) != 2 {
		t.Errorf("expect 2 children but %d, %+v", len(r.root.children), r.root)
	}

	err = r.Add(&metapb.API{
		ID:         3,
		URLPattern: "/(string):name",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error with const and string")
	}
	if len(r.root.children) != 3 {
		t.Errorf("expect 3 children but %d, %+v", len(r.root.children), r.root)
	}

	err = r.Add(&metapb.API{
		ID:         4,
		URLPattern: "/(number):age",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error with const, string, number")
	}
	if len(r.root.children) != 4 {
		t.Errorf("expect 4 children but %d, %+v", len(r.root.children), r.root)
	}

	err = r.Add(&metapb.API{
		ID:         5,
		URLPattern: "/(enum:off|on):action",
		Method:     "*",
	})
	if err != nil {
		t.Errorf("add error with const, string, number, enum")
	}
	if len(r.root.children) != 5 {
		t.Errorf("expect 4 children but %d, %+v", len(r.root.children), r.root)
	}

	r.Add(&metapb.API{
		ID:         6,
		URLPattern: "/a/b/b/a",
		Method:     "*",
	})
	r.Add(&metapb.API{
		ID:         7,
		URLPattern: "/a/b/b/c",
		Method:     "*",
	})
	_, ok := r.Find([]byte("/a/b/b/c"), "*", nil)
	if !ok {
		t.Errorf("expected match /a/b/b/c, but not")
	}
}

func TestFindWithStar(t *testing.T) {
	r := NewRoute()
	r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/*",
		Method:     "*",
	})

	r.Add(&metapb.API{
		ID:         2,
		URLPattern: "/users/1",
		Method:     "*",
	})

	params := make(map[string]string)
	id, _ := r.Find([]byte("/users/1"), "GET", func(name, value []byte) {
		params[string(name)] = string(value)
	})

	if id != 1 {
		t.Errorf("expect matched 1, but %d", id)
	}

	if params["*"] != "users/1" {
		t.Errorf("expect param value users/1, but %s", params["*"])
	}
}

func TestFindMatchAll(t *testing.T) {
	r := NewRoute()
	r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/*",
		Method:     "*",
	})

	params := make(map[string][]byte, 0)
	paramsFunc := func(name, value []byte) {
		params[string(name)] = value
	}

	id, _ := r.Find([]byte("/p1"), "GET", paramsFunc)
	if id != 1 {
		t.Errorf("expect matched 1, but %d", id)
	}
	if bytes.Compare(params["*"], []byte("p1")) != 0 {
		t.Errorf("expect params p1, but %s", params["*"])
	}

	id, _ = r.Find([]byte("/p1/p2"), "GET", paramsFunc)
	if id != 1 {
		t.Errorf("expect matched 1, but %d", id)
	}
	if bytes.Compare(params["*"], []byte("p1/p2")) != 0 {
		t.Errorf("expect params p1/p2, but %s", params["*"])
	}
}

func TestFind(t *testing.T) {
	r := NewRoute()
	r.Add(&metapb.API{
		ID:         1,
		URLPattern: "/",
		Method:     "*",
	})
	r.Add(&metapb.API{
		ID:         2,
		URLPattern: "/check",
		Method:     "*",
	})
	r.Add(&metapb.API{
		ID:         3,
		URLPattern: "/(string):name",
		Method:     "*",
	})
	r.Add(&metapb.API{
		ID:         4,
		URLPattern: "/(number):age",
		Method:     "*",
	})
	r.Add(&metapb.API{
		ID:         5,
		URLPattern: "/(enum:on|off):action",
		Method:     "*",
	})

	params := make(map[string][]byte, 0)
	paramsFunc := func(name, value []byte) {
		params[string(name)] = value
	}

	id, _ := r.Find([]byte("/"), "GET", paramsFunc)
	if id != 1 {
		t.Errorf("expect matched 1, but %d", id)
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/check"), "GET", paramsFunc)
	if id != 2 {
		t.Errorf("expect matched 2, but %d", id)
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/check2"), "GET", paramsFunc)
	if id != 3 {
		t.Errorf("expect matched 3, but %d", id)
	}
	if bytes.Compare(params["name"], []byte("check2")) != 0 {
		t.Errorf("expect params check2, but %s", params["name"])
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/123"), "GET", paramsFunc)
	if id != 4 {
		t.Errorf("expect matched 4, but %d", id)
	}
	if bytes.Compare(params["age"], []byte("123")) != 0 {
		t.Errorf("expect params 123, but %s", params["age"])
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/on"), "GET", paramsFunc)
	if id != 5 {
		t.Errorf("expect matched 5, but %d", id)
	}
	if bytes.Compare(params["action"], []byte("on")) != 0 {
		t.Errorf("expect params on, but %s", params["action"])
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/off"), "GET", paramsFunc)
	if id != 5 {
		t.Errorf("expect matched 5, but %d", id)
	}
	if bytes.Compare(params["action"], []byte("off")) != 0 {
		t.Errorf("expect params off, but %s", params["action"])
	}

	params = make(map[string][]byte, 0)
	id, _ = r.Find([]byte("/on/notmatches"), "GET", paramsFunc)
	if id != 0 {
		t.Errorf("expect not matched , but %d", id)
	}
}
