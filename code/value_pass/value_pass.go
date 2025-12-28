package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
)

type Status int

type Conference struct {
	ID                uuid.UUID
	ChairID           uuid.UUID
	Name              string
	Description       *string
	CoverID           *uuid.UUID
	LogoID            *uuid.UUID
	SubmissionStartAt time.Time
	SubmissionEndAt   time.Time
	ReviewStartAt     time.Time
	ReviewEndAt       time.Time
	ConferenceStartAt time.Time
	ConferenceEndAt   time.Time
	IsLectureRequired bool
	Status            Status
	CountryID         *int32
	StateID           *int32
	CityID            *int32
	Address           *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

const (
	runDuration = time.Minute
	batch       = 1_000_000
)

func main() {
	fmt.Println("===== VALUE PASSING TEST =====")

	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)

	start := time.Now()
	iterations := 0

	for time.Since(start) < runDuration {
		for i := 0; i < batch; i++ {
			c := newConference()
			layer1(c)
		}
		iterations += batch
	}

	runtime.ReadMemStats(&after)

	printStats(before, after, iterations)
}

func newConference() Conference {
	desc := "desc"
	addr := "addr"
	cid := int32(1)

	return Conference{
		ID:          uuid.New(),
		ChairID:     uuid.New(),
		Name:        "conf",
		Description: &desc,
		CountryID:   &cid,
		Address:     &addr,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// ---- value flow ----
func layer1(c Conference) { layer2(c) }
func layer2(c Conference) { layer3(c) }
func layer3(c Conference) { layer4(c) }
func layer4(c Conference) { _ = c.Name }

func printStats(b, a runtime.MemStats, it int) {
	fmt.Printf("Iterations: %d\n", it)
	fmt.Printf("TotalAlloc: %d MB\n", (a.TotalAlloc-b.TotalAlloc)/1024/1024)
	fmt.Printf("HeapAlloc:  %d MB\n", (a.HeapAlloc-b.HeapAlloc)/1024/1024)
	fmt.Printf("StackInuse: %d KB\n", (a.StackInuse-b.StackInuse)/1024)
	fmt.Printf("Mallocs:   %d\n", a.Mallocs-b.Mallocs)
	fmt.Printf("Frees:     %d\n", a.Frees-b.Frees)
	fmt.Printf("GC Count:  %d\n", a.NumGC-b.NumGC)
}
