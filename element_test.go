package conditions_test

import (
	"testing"
	"time"

	conditions "github.com/quan-to/selenium-go-conditions"
	"github.com/tebeka/selenium"
)

func TestElementIsLocated(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_add")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_add: %v\n", err)
	}

	// This should not raise an error.
	if err := wd.Wait(conditions.ElementIsLocated(selenium.ByID, "element")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err := wd.WaitWithTimeout(conditions.ElementIsLocated(selenium.ByID, "element2"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestElementIsVisible(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_show")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_show: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.ElementIsVisible(element)); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/element_show")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_show: %v\n", err)
	}
	element2, err := wd.FindElement(selenium.ByID, "element2")
	if err != nil {
		t.Fatalf("Cannot find #element2: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementIsVisible(element2), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestElementIsLocatedAndVisible(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_add_and_show")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_add_and_show: %v\n", err)
	}
	// This should not raise an error.
	if err = wd.Wait(conditions.ElementIsLocatedAndVisible(selenium.ByID, "element")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/element_add_and_show")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_add_and_show: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementIsLocatedAndVisible(selenium.ByID, "element2"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestElementIsEnabled(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_enable")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_enable: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.ElementIsEnabled(element)); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/element_enable")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_enable: %v\n", err)
	}
	element2, err := wd.FindElement(selenium.ByID, "element2")
	if err != nil {
		t.Fatalf("Cannot find #element2: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementIsEnabled(element2), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestElementTextIs(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_change_text")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_change_text: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.ElementTextIs(element, "Text changed.")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}
	element, err = wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementTextIs(element, "Another page."), 500*time.Millisecond); err == nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}
}

func TestElementTextContains(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_change_text")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_change_text: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.ElementTextContains(element, "changed")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}
	element, err = wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementTextContains(element, "Another"), 500*time.Millisecond); err == nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}
}

func TestElementAttributeIs(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/element_change_attribute")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/element_change_attribute: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.ElementAttributeIs(element, "data", "Some data.")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}
	element, err = wd.FindElement(selenium.ByID, "element")
	if err != nil {
		t.Fatalf("Cannot find #element: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.ElementAttributeIs(element, "data", "Another data."), 500*time.Millisecond); err == nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}
}
