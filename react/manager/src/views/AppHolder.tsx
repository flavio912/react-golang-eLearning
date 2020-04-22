import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";

type Props = {
  children?: React.ReactChildren;
};

export const AppHolder = ({ children }: Props) => {
  const tabs: Tab[] = [
    {
      id: 0,
      icon: "TTC_Logo_Icon",
      children: <div></div>,
    },
  ];
  return (
    <div>
      <HeaderMenu logo={""} user={{ name: "Test", url: "My URL" }} />
      <SideMenu tabs={tabs} selected={tabs[0]} onClick={() => {}} />
      {children}
    </div>
  );
};
