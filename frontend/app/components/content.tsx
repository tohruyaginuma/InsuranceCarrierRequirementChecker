import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<{}>;

const Content = (props: Props) => {
  const { children } = props;
  return <div className="w-full mt-4">{children}</div>;
};

export default Content;
