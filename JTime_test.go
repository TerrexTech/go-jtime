package jtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

type DateTest struct {
	DateArrived JTime `bson:"date_arrived,omitempty" json:"date_arrived,omitempty"`
}

var _ = Describe("Test custom time", func() {
	var (
		testJSON string
		timeTest *DateTest
	)

	BeforeEach(func() {
		timeTest = &DateTest{}
		testJSON = `{"date_arrived": "2018-09-13T00:32:23.534Z"}`
	})

	It("Should unmarshal json into DateTest", func() {
		err := json.Unmarshal([]byte(testJSON), &timeTest)
		Expect(err).ToNot(HaveOccurred())

		expectedTime, err := time.Parse(
			"2006-01-02 15:04:05 -0700 MST",
			"2018-09-13 00:32:23.999 +0000 UTC",
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(expectedTime.Unix()).To(Equal(timeTest.DateArrived.Unix()))
	})

	It("Should marshal DateTest into json", func() {
		t := time.Now()
		jt := &JTime{t}

		mtime, err := json.Marshal(jt)
		expectedTime := t.Format(time.RFC3339Nano)
		expectedTimeJSON := fmt.Sprintf(`"%s"`, expectedTime)

		Expect(err).ToNot(HaveOccurred())
		Expect(string(mtime)).To(Equal(expectedTimeJSON))
	})
})
