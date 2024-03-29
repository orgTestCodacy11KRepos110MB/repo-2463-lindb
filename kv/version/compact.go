// Licensed to LinDB under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. LinDB licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package version

import "github.com/lindb/lindb/kv/table"

// Compaction represents the compaction job context.
type Compaction struct {
	level int

	inputs        [][]*FileMeta
	levelInputs   []*FileMeta
	levelUpInputs []*FileMeta

	editLog EditLog
}

// NewCompaction create a compaction job context.
func NewCompaction(familyID FamilyID, level int, levelInputs, levelUpInputs []*FileMeta) *Compaction {
	return &Compaction{
		level:         level,
		inputs:        [][]*FileMeta{levelInputs, levelUpInputs},
		levelInputs:   levelInputs,
		levelUpInputs: levelUpInputs,
		editLog:       NewEditLog(familyID),
	}
}

// IsTrivialMove returns a trivial compaction that can be implemented by just
// moving a single input file to the next level (no merging or splitting).
// returns true: can just move file to the next level.
func (c *Compaction) IsTrivialMove() bool {
	return len(c.levelInputs) == 1 && len(c.levelUpInputs) == 0
}

// GetLevelFiles returns low level files.
func (c *Compaction) GetLevelFiles() []*FileMeta {
	return c.levelInputs
}

// DeleteFile deletes an old file which compaction input file.
func (c *Compaction) DeleteFile(level int, fileNumber table.FileNumber) {
	c.editLog.Add(NewDeleteFile(int32(level), fileNumber))
}

// AddFile adds a new file which compaction output file.
func (c *Compaction) AddFile(level int, file *FileMeta) {
	c.editLog.Add(CreateNewFile(int32(level), file))
}

// AddReferenceFiles adds new reference file for checking if file already rollup.
func (c *Compaction) AddReferenceFiles(newReferenceFileLogs []Log) {
	for _, log := range newReferenceFileLogs {
		c.editLog.Add(log)
	}
}

// MarkInputDeletes marks all inputs of this compaction as deletion, adds log to version edit.
func (c *Compaction) MarkInputDeletes() {
	for _, input := range c.levelInputs {
		c.editLog.Add(NewDeleteFile(int32(c.level), input.fileNumber))
	}
	for _, upInput := range c.levelUpInputs {
		c.editLog.Add(NewDeleteFile(int32(c.level+1), upInput.fileNumber))
	}
}

// GetLevel returns compaction level.
func (c *Compaction) GetLevel() int {
	return c.level
}

// GetEditLog returns edit log.
func (c *Compaction) GetEditLog() EditLog {
	return c.editLog
}

// GetInputs returns all input files for compaction job.
func (c *Compaction) GetInputs() [][]*FileMeta {
	return c.inputs
}
