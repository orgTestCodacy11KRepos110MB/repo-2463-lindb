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
import { Layout, Nav, Space, Tag } from "@douyinfe/semi-ui";
import Logo from "@src/assets/logo_dark.svg";
import { defaultOpenKeys, menus, routeMap } from "@src/configs";
import { URLStore } from "@src/stores";
import * as _ from "lodash-es";
import { reaction } from "mobx";
import React, { useEffect, useState } from "react";
import { useWatchURLChange } from "@src/hooks";
const { Sider } = Layout;

export type SiderMenuProps = {
  defaultOpenAll?: boolean;
};

export default function SiderMenu(props: SiderMenuProps) {
  const { defaultOpenAll } = props;
  const [selectedKeys, setSelectedKeys] = useState([] as string[]);

  useWatchURLChange(() => {
    const selectedKeys: string[] = [];
    const path = URLStore.path;
    const findSelectedKeys = (menus: any[]) => {
      (menus || []).map((item) => {
        if (path.includes(item.path)) {
          selectedKeys.push(item.path);
        }
        if (item.items) {
          findSelectedKeys(item.items);
        }
      });
    };
    findSelectedKeys(menus);
    setSelectedKeys(selectedKeys);
  });

  return (
    <Sider>
      <Nav
        defaultOpenKeys={defaultOpenAll ? defaultOpenKeys : []}
        style={{
          maxWidth: 220,
          height: "100%",
        }}
        items={menus as any[]}
        selectedKeys={selectedKeys}
        onClick={(data) => {
          const item = routeMap.get(`${data.itemKey}`);
          const needClearKeys = _.pullAll(
            URLStore.getParamKeys(),
            _.get(item, "keep", [])
          );
          console.log(
            "mmmmmm",
            item?.path,
            needClearKeys,
            Object.keys(URLStore.params)
          );
          URLStore.changeURLParams({
            path: item?.path,
            needDelete: needClearKeys,
          });
        }}
        header={{
          logo: <img src={Logo} style={{ width: 48, height: 48 }} />,
          text: (
            <Space align="end">
              <div
                style={{
                  fontSize: 32,
                  height: 32,
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                }}
              >
                LinDB
              </div>
              <Tag type="light" color="yellow" size="small">
                v0.1
              </Tag>
            </Space>
          ),
          style: { paddingTop: 12, paddingBottom: 12, paddingLeft: 2 },
        }}
      />
    </Sider>
  );
}