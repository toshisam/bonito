package main

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datetime functions", func() {
	Context("ParseTime", func() {
		It("should return now for now", func() {
			Expect(ParseTime("now")).To(BeTemporally("~", time.Now()))
		})

		It("should work for hours", func() {
			Expect(ParseTime("now-1h")).To(BeTemporally("~", time.Now().Add(-1*time.Hour)))
		})

		It("should work for seconds", func() {
			Expect(ParseTime("now-4s")).To(BeTemporally("~", time.Now().Add(-4*time.Second)))
		})

		It("should work for minutes", func() {
			Expect(ParseTime("now+6m")).To(BeTemporally("~", time.Now().Add(6*time.Minute)))
		})

		It("should work for days", func() {
			Expect(ParseTime("now-16d")).To(BeTemporally("~", time.Now().AddDate(0, 0, -16)))
		})

		It("should work for weeks", func() {
			Expect(ParseTime("now+2w")).To(BeTemporally("~", time.Now().AddDate(0, 0, 14)))
		})

		It("should work for months", func() {
			Expect(ParseTime("now-2M")).To(BeTemporally("~", time.Now().AddDate(0, -2, 0)))
		})

		It("should work for years", func() {
			Expect(ParseTime("now-1y")).To(BeTemporally("~", time.Now().AddDate(-1, 0, 0)))
		})

		It("should work for absolute date", func() {
			Expect(ParseTime("2015-01-24T14:06:05.071Z")).To(BeTemporally("~",
				time.Date(2015, time.January, 24, 14, 06, 05, 71*1e6, time.UTC)))
		})

		It("should fail for invalid spec", func() {
			_, err := ParseTime("then+4h")
			Expect(err).To(HaveOccurred())
		})

		It("should fail when the value is not specified", func() {
			_, err := ParseTime("now-h")
			Expect(err).To(HaveOccurred())
		})

		It("should fail when microseconds are given", func() {
			_, err := ParseTime("2015-01-24T14:06:05.071000Z")
			Expect(err).To(HaveOccurred())
		})
	})
})
