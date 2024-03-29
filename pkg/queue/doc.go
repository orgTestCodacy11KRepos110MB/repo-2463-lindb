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

package queue

/**
FanOutQueue structure
+---------------------------------------------------------+
| Tail             FanOutQueue                    Head    |
|                                                         |
|                                                         |
|    +-----------+   +-----------+                        |
|    |  Segment  |   |  Segment..|                        |
|    |  Begin    |   |           |                        |
|    |  End      |   |           |                        |
|    |  Append   |   |           |                        |
|    |  Read     |   |           |                        |
|    |           |   |           |                        |
|    +----+------+   +------+----+                        |
|         ^                 ^                             |
+---------+-----------------+-----------------------------+
|         |                 |                             ^
|         |                 |                             |
+---------+--+   +----------+-+                           | Append
|  ConsumerGroup    |   |   ConsumerGroup.. |                           |
|  Name      |   |            |                           + +--------+
|  Consume   |   |            |                             | Data...|
|  Get       |   |            |                             |        |
|  Ack       |   |            |                             |        |
|            |   |            |                             +--------+
+------------+   +------------+

Segment structure
+-----------------------------------------------+
|            Segment                            |
|                                               |
|    +----------------------+    +----------+   |
|    |    IndexPage         |    | DataPage |   |
|    |  dataOffset dataLen+------->Message1 |   |
|    |  (4 bytes) (4 bytes) |    | Message2 |   |
|    |  ....                |    | ....     |   |
|    |                      |    |          |   |
|    |                      |    |          |   |
|    +----------------------+    +----------+   |
|                                               |
+-----------------------------------------------+


directory structure
/fanoutQueueDir
	queue.meta
	/segments
		0.idx
		0.dat
		100.idx
		100.dat
		...
	/consumerGroup
		fanOutName1.meta
		fanOutName2.meta
		...
*/
