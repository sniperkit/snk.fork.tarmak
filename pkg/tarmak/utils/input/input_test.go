/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package input_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/sniperkit/snk.fork.tarmak/pkg/tarmak/utils/input"
)

func Test_Input_AskYesNo_Yes(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskYesNo{
		Query:   "Should this test pass?",
		Default: true,
	}

	go func() {
		inWriter.Write([]byte("Y\n"))
		//inWriter.Close()
	}()

	resp, err := i.AskYesNo(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if !resp {
		t.Error("expected true response")
	}
}

func Test_Input_AskYesNo_No(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskYesNo{
		Query:   "Should this test pass?",
		Default: true,
	}

	go func() {
		inWriter.Write([]byte("no\n"))
		//inWriter.Close()
	}()

	resp, err := i.AskYesNo(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if resp {
		t.Error("expected false response")
	}
}

func Test_Input_AskYesNo_Default_No(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskYesNo{
		Query:   "Should this test pass?",
		Default: false,
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskYesNo(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if resp {
		t.Error("expected false response")
	}
}

func Test_Input_AskYesNo_Wrong_Reask_Yes(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskYesNo{
		Query:   "Should this test pass?",
		Default: false,
	}

	go func() {
		inWriter.Write([]byte("xxx\n"))
		inWriter.Write([]byte("Y\n"))
	}()

	resp, err := i.AskYesNo(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if !resp {
		t.Error("expected true response")
	}
}

func Test_Input_AskOpen_NoDefault_NoResponse_DisallowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query: "Should this test pass?",
	}

	go func() {
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("no_default\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no_default", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_NoDefault_NoResponse_AllowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:      "Should this test pass?",
		AllowEmpty: true,
	}

	go func() {
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("no_default\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if resp != "" {
		t.Errorf("unexpected response, exp=<nothing> act=%s", resp)
	}
}

func Test_Input_AskOpen_Default_NoResponse_DisallowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:   "Should this test pass?",
		Default: "no never",
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no never", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_Default_NoResponse_AllowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:      "Should this test pass?",
		Default:    "no never",
		AllowEmpty: true,
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no never", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_NoDefault_Response_DisallowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query: "Should this test pass?",
	}

	go func() {
		inWriter.Write([]byte("no_default\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no_default", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_NoDefault_Response_AllowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:      "Should this test pass?",
		AllowEmpty: true,
	}

	go func() {
		inWriter.Write([]byte("no_default\n"))
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no_default", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_Default_Response_DisallowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:   "Should this test pass?",
		Default: "no never",
	}

	go func() {
		inWriter.Write([]byte("no_default\n"))
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no_default", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskOpen_Default_Response_AllowEmpty(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskOpen{
		Query:      "Should this test pass?",
		Default:    "no never",
		AllowEmpty: true,
	}

	go func() {
		inWriter.Write([]byte("no_default\n"))
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskOpen(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := "no_default", resp; exp != act {
		t.Errorf("unexpected response, exp=%s act=%s", exp, act)
	}
}

func Test_Input_AskSelection_NoDefault_NoResponse(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskSelection{
		Query:   "Should this test pass?",
		Choices: []string{"choice1", "choice2", "choice3"},
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := 0, resp; exp != act {
		t.Errorf("unexpected response, exp=%d act=%d", exp, act)
	}
}

func Test_Input_AskSelection_Default_NoResponse(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskSelection{
		Query:   "Should this test pass?",
		Choices: []string{"choice1", "choice2", "choice3"},
		Default: 2,
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := 2, resp; exp != act {
		t.Errorf("unexpected response, exp=%d act=%d", exp, act)
	}
}

func Test_Input_AskSelection_NoDefault_Response(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskSelection{
		Query:   "Should this test pass?",
		Choices: []string{"choice1", "choice2", "choice3"},
	}

	go func() {
		inWriter.Write([]byte("foo\n"))
		inWriter.Write([]byte("2bar32\n"))
		inWriter.Write([]byte("2\n"))
	}()

	resp, err := i.AskSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := 1, resp; exp != act {
		t.Errorf("unexpected response, exp=%d act=%d", exp, act)
	}
}

func Test_Input_AskSelection_Default_Response(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskSelection{
		Query:   "Should this test pass?",
		Choices: []string{"choice1", "choice2", "choice3"},
		Default: 3,
	}

	go func() {
		inWriter.Write([]byte("foo\n"))
		inWriter.Write([]byte("2bar32\n"))
		inWriter.Write([]byte("2\n"))
	}()

	resp, err := i.AskSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := 1, resp; exp != act {
		t.Errorf("unexpected response, exp=%d act=%d", exp, act)
	}
}

func Test_Input_AskMultiSelection_NoSelected_Min2_Max3_EnterCont(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskMultipleSelection{
		AskSelection: &input.AskSelection{
			Query:   "Should this pass?",
			Choices: []string{"1", "2", "3", "4"},
			Default: 1,
		},
		SelectedChoices: []bool{false, false, false, false},
		MinSelected:     2,
		MaxSelected:     3,
	}

	go func() {
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("1\n"))
		inWriter.Write([]byte("2\n"))
		inWriter.Write([]byte("3\n"))
		inWriter.Write([]byte("1foo\n"))
		inWriter.Write([]byte("1bar\n"))
		inWriter.Write([]byte("4\n"))
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("4\n"))
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskMultipleSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := []string{"1", "2", "3"}, resp; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected response, exp=%s act=%s", exp, resp)
	}
}

func Test_Input_AskMultiSelection_Selected_Min2_Max3_EnterCont(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskMultipleSelection{
		AskSelection: &input.AskSelection{
			Query:   "Should this pass?",
			Choices: []string{"1", "2", "3", "4"},
			Default: 1,
		},
		SelectedChoices: []bool{true, true, false, false},
		MinSelected:     2,
		MaxSelected:     3,
	}

	go func() {
		inWriter.Write([]byte("\n"))
	}()

	resp, err := i.AskMultipleSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := []string{"1", "2"}, resp; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected response, exp=%s act=%s", exp, resp)
	}
}

func Test_Input_AskMultiSelection_NoSelected_Min2_Max3_NumCont(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskMultipleSelection{
		AskSelection: &input.AskSelection{
			Query:   "Should this pass?",
			Choices: []string{"1", "2", "3", "4"},
			Default: 1,
		},
		SelectedChoices: []bool{false, false, false, false},
		MinSelected:     2,
		MaxSelected:     3,
	}

	go func() {
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("1foo\n"))
		inWriter.Write([]byte("1bar\n"))
		inWriter.Write([]byte("1\n"))
		inWriter.Write([]byte("2\n"))
		inWriter.Write([]byte("3\n"))
		inWriter.Write([]byte("4\n"))
		inWriter.Write([]byte("\n"))
		inWriter.Write([]byte("4\n"))
		inWriter.Write([]byte("5\n"))
	}()

	resp, err := i.AskMultipleSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := []string{"1", "2", "3"}, resp; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected response, exp=%s act=%s", exp, resp)
	}
}

func Test_Input_AskMultiSelection_Selected_Min2_Max3_NumCont(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskMultipleSelection{
		AskSelection: &input.AskSelection{
			Query:   "Should this pass?",
			Choices: []string{"1", "2", "3", "4"},
			Default: 1,
		},
		SelectedChoices: []bool{true, true, false, false},
		MinSelected:     2,
		MaxSelected:     3,
	}

	go func() {
		inWriter.Write([]byte("1foo\n"))
		inWriter.Write([]byte("1bar\n"))
		inWriter.Write([]byte("5\n"))
	}()

	resp, err := i.AskMultipleSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := []string{"1", "2"}, resp; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected response, exp=%s act=%s", exp, resp)
	}
}

func Test_Input_AskMultiSelection_NoSelected_Min2_Max3_NumCont_Toggle(t *testing.T) {
	out := new(bytes.Buffer)
	inReader, inWriter := io.Pipe()
	i := input.New(inReader, out)

	question := &input.AskMultipleSelection{
		AskSelection: &input.AskSelection{
			Query:   "Should this pass?",
			Choices: []string{"1", "2", "3", "4"},
			Default: 1,
		},
		SelectedChoices: []bool{false, false, false, false},
		MinSelected:     2,
		MaxSelected:     3,
	}

	go func() {
		inWriter.Write([]byte("1foo\n"))
		inWriter.Write([]byte("1bar\n"))
		inWriter.Write([]byte("5\n"))
		inWriter.Write([]byte("1\n"))
		inWriter.Write([]byte("2\n"))
		inWriter.Write([]byte("1\n"))
		inWriter.Write([]byte("2\n"))
		inWriter.Write([]byte("5\n"))
		inWriter.Write([]byte("1\n"))
		inWriter.Write([]byte("2\n"))
		inWriter.Write([]byte("5\n"))
	}()

	resp, err := i.AskMultipleSelection(question)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if exp, act := []string{"1", "2"}, resp; !reflect.DeepEqual(exp, act) {
		t.Errorf("unexpected response, exp=%s act=%s", exp, resp)
	}
}
