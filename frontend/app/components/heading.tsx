import type { PropsWithChildren } from "react";

type Props = PropsWithChildren<{
  title: string;
}>;

const Heading = (props: Props) => {
  const { title } = props;
  return (
    <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight text-balance">
      {title}
    </h1>
  );
};

export default Heading;
