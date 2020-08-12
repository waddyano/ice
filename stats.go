//  Copyright (c) 2020 Bluge Labs, LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ice

import (
	segment "github.com/blugelabs/bluge_segment_api"
)

type CollectionStats struct {
	totalDocCount    int
	docCount         int
	sumTotalTermFreq int
}

func (c *CollectionStats) TotalDocumentCount() int {
	return c.totalDocCount
}

func (c *CollectionStats) DocumentCount() int {
	return c.docCount
}

func (c *CollectionStats) SumTotalTermFrequency() int {
	return c.sumTotalTermFreq
}

func (c *CollectionStats) Merge(other segment.CollectionStats) {
	c.totalDocCount += other.TotalDocumentCount()
	c.docCount += other.DocumentCount()
	c.sumTotalTermFreq += other.SumTotalTermFrequency()
}

func (s *Segment) CollectionStats(field string) (segment.CollectionStats, error) {
	var rv = &CollectionStats{}
	fieldIDPlus1 := s.fieldsMap[field]
	if fieldIDPlus1 > 0 {
		rv.totalDocCount = int(s.footer.numDocs)
		rv.docCount = s.fieldDocs[fieldIDPlus1-1]
		rv.sumTotalTermFreq = s.fieldFreqs[fieldIDPlus1-1]
	}
	return rv, nil
}
