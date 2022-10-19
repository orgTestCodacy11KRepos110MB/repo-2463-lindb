/*
Licensed to LinDB under one or more contributor
license agreements. See the NOTICE file distributed with
this work for additional information regarding copyright
ownership. LinDB licenses this file to you under
the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
 
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

import { IconCrossCircleStroked } from "@douyinfe/semi-icons";
import { Empty, Spin, Tooltip, Typography } from "@douyinfe/semi-ui";
import { ApiKit } from "@src/utils";
import React, { CSSProperties } from "react";
import { Icon } from "@src/components";

const { Text } = Typography;

const SimpleStatusTip: React.FC<{
  isLoading: boolean;
  isError: boolean;
  error: any;
  style?: CSSProperties;
}> = (props) => {
  const { isLoading, isError, error, style } = props;
  if (isLoading) {
    return <Spin style={{ width: 18, height: 18 }} />;
  }
  if (isError) {
    if (ApiKit.getErrorCode(error) === 404) {
      return (
        <Empty image={<Icon icon="iconempty" style={{ fontSize: 16 }} />} />
      );
    }
    return (
      <Text type="danger">
        <Tooltip content={ApiKit.getErrorMsg(error)}>
          <IconCrossCircleStroked />
        </Tooltip>
      </Text>
    );
  }
  return null;
};

export default SimpleStatusTip;
