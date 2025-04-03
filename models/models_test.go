package models

import (
	"context"
	"reflect"
	"testing"
)

func TestSearchIndex_PreloadData(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		data []Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Single artist with members and locations",
			fields: fields{
				ArtistName:   make(map[string][]IndexedData),
				MemberName:   make(map[string][]IndexedData),
				LocationName: make(map[string][]IndexedData),
				FirstAlbum:   make(map[string][]IndexedData),
				CreationDate: make(map[int][]IndexedData),
			},
			args: args{
				data: []Data{
					{
						A: Artist{
							Name:         "Band A",
							Members:      []string{"Member 1", "Member 2"},
							FirstAlbum:   "First Album",
							CreationDate: 1990,
						},
						L: Location{
							Locations: []string{"New York", "London"},
						},
					},
				},
			},
			want: fields{
				ArtistName: map[string][]IndexedData{
					"band a": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
				},
				MemberName: map[string][]IndexedData{
					"member 1": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
					"member 2": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
				},
				LocationName: map[string][]IndexedData{
					"new york": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
					"london": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
				},
				FirstAlbum: map[string][]IndexedData{
					"first album": {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
				},
				CreationDate: map[int][]IndexedData{
					1990: {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									Members:      []string{"Member 1", "Member 2"},
									FirstAlbum:   "First Album",
									CreationDate: 1990,
								},
								L: Location{
									Locations: []string{"New York", "London"},
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				ArtistName:   tt.fields.ArtistName,
				MemberName:   tt.fields.MemberName,
				LocationName: tt.fields.LocationName,
				FirstAlbum:   tt.fields.FirstAlbum,
				CreationDate: tt.fields.CreationDate,
			}
			index.PreloadData(tt.args.data)

			// Compare actual results with expected "want" results.
			if !reflect.DeepEqual(index.ArtistName, tt.want.ArtistName) {
				t.Errorf("ArtistName = %v, want %v", index.ArtistName, tt.want.ArtistName)
			}
			if !reflect.DeepEqual(index.MemberName, tt.want.MemberName) {
				t.Errorf("MemberName = %v, want %v", index.MemberName, tt.want.MemberName)
			}
			if !reflect.DeepEqual(index.LocationName, tt.want.LocationName) {
				t.Errorf("LocationName = %v, want %v", index.LocationName, tt.want.LocationName)
			}
			if !reflect.DeepEqual(index.FirstAlbum, tt.want.FirstAlbum) {
				t.Errorf("FirstAlbum = %v, want %v", index.FirstAlbum, tt.want.FirstAlbum)
			}
			if !reflect.DeepEqual(index.CreationDate, tt.want.CreationDate) {
				t.Errorf("CreationDate = %v, want %v", index.CreationDate, tt.want.CreationDate)
			}
		})
	}
}

func TestSearchIndex_searchCreationDates(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		ctx        context.Context
		query      string
		resultChan chan SearchResult
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult
	}{
		{
			name: "Valid creation date found",
			fields: fields{
				CreationDate: map[int][]IndexedData{
					1990: {
						{
							Data: Data{
								A: Artist{
									Name:         "Band A",
									CreationDate: 1990,
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "1990",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{
				{
					Result:     "1990 - creation date-Band A",
					ArtistName: "Band A",
				},
			},
		},
		{
			name: "Creation date not found",
			fields: fields{
				CreationDate: map[int][]IndexedData{
					1985: {
						{
							Data: Data{
								A: Artist{
									Name:         "Band B",
									CreationDate: 1985,
								},
							},
							ArtistName: "Band B",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "1990",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				CreationDate: tt.fields.CreationDate,
			}

			results := []SearchResult{}
			done := make(chan struct{})
			go func() {
				for result := range tt.args.resultChan {
					results = append(results, result)
				}
				close(done)
			}()

			index.searchCreationDates(tt.args.ctx, tt.args.query, tt.args.resultChan)
			close(tt.args.resultChan)

			<-done

			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("searchCreationDates() results = %v, want %v", results, tt.want)
			}
		})
	}
}

func TestSearchIndex_searchFirstAlbums(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
		// mu           sync.RWMutex
	}
	type args struct {
		ctx        context.Context
		query      string
		resultChan chan SearchResult // Bidirectional channel for testing
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult // Expected results
	}{
		{
			name: "First album match found",
			fields: fields{
				FirstAlbum: map[string][]IndexedData{
					"the album": {
						{
							Data: Data{
								A: Artist{
									Name:       "Band A",
									FirstAlbum: "The Album",
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "the album",
				resultChan: make(chan SearchResult, 1), // Buffered channel to avoid blocking
			},
			want: []SearchResult{
				{
					Result:     "The Album - first album-Band A",
					ArtistName: "Band A",
				},
			},
		},
		{
			name: "No first album match",
			fields: fields{
				FirstAlbum: map[string][]IndexedData{
					"another album": {
						{
							Data: Data{
								A: Artist{
									Name:       "Band B",
									FirstAlbum: "Another Album",
								},
							},
							ArtistName: "Band B",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "not found",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No matches
		},
		{
			name: "Context cancellation",
			fields: fields{
				FirstAlbum: map[string][]IndexedData{
					"the album": {
						{
							Data: Data{
								A: Artist{
									Name:       "Band A",
									FirstAlbum: "The Album",
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel() // Immediately cancel the context
					return ctx
				}(),
				query:      "the album",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No results due to context cancellation
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				FirstAlbum: tt.fields.FirstAlbum,
				// mu:         tt.fields.mu,
			}

			// Collect results asynchronously
			results := []SearchResult{}
			done := make(chan struct{})
			go func() {
				for result := range tt.args.resultChan {
					results = append(results, result)
				}
				close(done)
			}()

			// Call the searchFirstAlbums function
			index.searchFirstAlbums(tt.args.ctx, tt.args.query, tt.args.resultChan)
			close(tt.args.resultChan) // Close channel after sending results

			<-done // Wait for the collection goroutine to finish

			// Verify the results
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("searchFirstAlbums() = %v, want %v", results, tt.want)
			}
		})
	}
}

func TestSearchIndex_searchLocations(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		ctx        context.Context
		query      string
		resultChan chan SearchResult // Bidirectional channel for testing
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult // Expected results
	}{
		{
			name: "Location match found",
			fields: fields{
				LocationName: map[string][]IndexedData{
					"new york": {
						{
							Data: Data{
								L: Location{
									Locations: []string{"New York"},
								},
								A: Artist{
									Name: "Band A",
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "new york",
				resultChan: make(chan SearchResult, 1), // Buffered channel to avoid blocking
			},
			want: []SearchResult{
				{
					Result:     "New York - location-Band A",
					ArtistName: "Band A",
				},
			},
		},
		{
			name: "No location match",
			fields: fields{
				LocationName: map[string][]IndexedData{
					"los angeles": {
						{
							Data: Data{
								L: Location{
									Locations: []string{"Los Angeles"},
								},
								A: Artist{
									Name: "Band B",
								},
							},
							ArtistName: "Band B",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "chicago",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No matches
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				LocationName: tt.fields.LocationName,
				// mu:           tt.fields.mu,
			}

			// Collect results asynchronously
			results := []SearchResult{}
			done := make(chan struct{})
			go func() {
				for result := range tt.args.resultChan {
					results = append(results, result)
				}
				close(done)
			}()

			// Call the searchLocations function
			index.searchLocations(tt.args.ctx, tt.args.query, tt.args.resultChan)
			close(tt.args.resultChan) // Close channel after sending results

			<-done // Wait for the collection goroutine to finish

			// Verify the results
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("searchLocations() = %v, want %v", results, tt.want)
			}
		})
	}
}

func TestSearchIndex_searchMemberNames(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		ctx        context.Context
		query      string
		resultChan chan SearchResult
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult
	}{
		{
			name: "Member name match found",
			fields: fields{
				MemberName: map[string][]IndexedData{
					"john doe": {
						{
							Data: Data{
								A: Artist{
									Members: []string{"John Doe", "Jane Smith"},
									Name:    "Band A",
								},
							},
							ArtistName: "Band A",
						},
					},
				},
			},

			args: args{
				ctx:        context.Background(),
				query:      "john",
				resultChan: make(chan SearchResult, 1), // Buffered channel to avoid blocking
			},
			want: []SearchResult{
				{
					Result:     "John Doe - member of Band A band",
					ArtistName: "Band A",
				},
			},
		},
		{
			name: "No member name match",
			fields: fields{
				MemberName: map[string][]IndexedData{
					"jane smith": {
						{
							Data: Data{
								A: Artist{
									Members: []string{"Jane Smith", "Alice Johnson"},
									Name:    "Band B",
								},
							},
							ArtistName: "Band B",
						},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "bob",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No matches
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				MemberName: tt.fields.MemberName,
			}

			// Collect results asynchronously
			results := []SearchResult{}
			done := make(chan struct{})

			go func() {
				for result := range tt.args.resultChan {
					results = append(results, result)
				}
				close(done)
			}()

			// Call the searchMemberNames function
			index.searchMemberNames(tt.args.ctx, tt.args.query, tt.args.resultChan)
			close(tt.args.resultChan) // Close channel after sending results

			<-done // Wait for the collection goroutine to finish

			// Verify the results
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("searchMemberNames() = %v, want %v", results, tt.want)
			}
		})
	}
}

func TestSearchIndex_searchArtistNames(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		ctx        context.Context
		query      string
		resultChan chan SearchResult
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult
	}{
		{
			name: "Artist name match found",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"the beatles": {
						{Data: Data{
							A: Artist{Name: "The Beatles"},
						}, ArtistName: "The Beatles"},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "beatles",
				resultChan: make(chan SearchResult, 1), // Buffered channel
			},
			want: []SearchResult{
				{Result: "The Beatles - artist/band", ArtistName: "The Beatles"},
			},
		},
		{
			name: "No artist name match",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"pink floyd": {
						{Data: Data{
							A: Artist{Name: "Pink Floyd"},
						}, ArtistName: "Pink Floyd"},
					},
				},
			},
			args: args{
				ctx:        context.Background(),
				query:      "radiohead",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No matches
		},
		{
			name: "Context cancellation",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"led zeppelin": {
						{Data: Data{
							A: Artist{Name: "Led Zeppelin"},
						}, ArtistName: "Led Zeppelin"},
					},
				},
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithCancel(context.Background())
					cancel() // Immediately cancel the context
					return ctx
				}(),
				query:      "led",
				resultChan: make(chan SearchResult, 1),
			},
			want: []SearchResult{}, // No results due to context cancellation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				ArtistName:   tt.fields.ArtistName,
				MemberName:   tt.fields.MemberName,
				LocationName: tt.fields.LocationName,
				FirstAlbum:   tt.fields.FirstAlbum,
				CreationDate: tt.fields.CreationDate,
			}

			// Collect results asynchronously
			results := []SearchResult{}
			done := make(chan struct{})

			go func() {
				for result := range tt.args.resultChan {
					results = append(results, result)
				}
				close(done)
			}()

			// Call the searchArtistNames function
			index.searchArtistNames(tt.args.ctx, tt.args.query, tt.args.resultChan)
			close(tt.args.resultChan) // Close channel after sending results

			<-done // Wait for the collection goroutine to finish

			// Verify the results
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("searchArtistNames() = %v, want %v", results, tt.want)
			}
		})
	}
}

func TestSearchIndex_Search(t *testing.T) {
	type fields struct {
		ArtistName   map[string][]IndexedData
		MemberName   map[string][]IndexedData
		LocationName map[string][]IndexedData
		FirstAlbum   map[string][]IndexedData
		CreationDate map[int][]IndexedData
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []SearchResult
	}{
		{
			name: "Artist name match found",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"the beatles": {
						{Data: Data{
							A: Artist{Name: "The Beatles"},
						}, ArtistName: "The Beatles"},
					},
				},
			},
			args: args{
				query: "beatles",
			},
			want: []SearchResult{
				{Result: "The Beatles - artist/band", ArtistName: "The Beatles"},
			},
		},
		{
			name: "Multiple matches across categories",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"pink floyd": {
						{Data: Data{
							A: Artist{Name: "Pink Floyd"},
						}, ArtistName: "Pink Floyd"},
					},
				},
				MemberName: map[string][]IndexedData{
					"roger waters": {
						{Data: Data{
							A: Artist{Members: []string{"Roger Waters"}},
						}, ArtistName: "Pink Floyd"},
					},
				},
				LocationName: map[string][]IndexedData{
					"london": {
						{Data: Data{
							L: Location{Locations: []string{"London"}},
						}, ArtistName: "Pink Floyd"},
					},
				},
			},
			args: args{
				query: "pink",
			},
			want: []SearchResult{
				{Result: "Pink Floyd - artist/band", ArtistName: "Pink Floyd"},
			},
		},
		{
			name: "No matches",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"the who": {
						{Data: Data{
							A: Artist{Name: "The Who"},
						}, ArtistName: "The Who"},
					},
				},
			},
			args: args{
				query: "u2",
			},
			want: []SearchResult{}, // No matches
		},
		{
			name: "Empty query returns no results",
			fields: fields{
				ArtistName: map[string][]IndexedData{
					"the who": {
						{Data: Data{
							A: Artist{Name: "The Who"},
						}, ArtistName: "The Who"},
					},
				},
			},
			args: args{
				query: "",
			},
			want: []SearchResult{}, // No matches
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := &SearchIndex{
				ArtistName:   tt.fields.ArtistName,
				MemberName:   tt.fields.MemberName,
				LocationName: tt.fields.LocationName,
				FirstAlbum:   tt.fields.FirstAlbum,
				CreationDate: tt.fields.CreationDate,
			}
			if got := index.Search(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchIndex.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
