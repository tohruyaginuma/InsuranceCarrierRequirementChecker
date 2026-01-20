import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<{}>;

const Content = (props: Props) => {
	const { children } = props;
	return <div className="mt-6">{children}</div>;
};

export default Content;
