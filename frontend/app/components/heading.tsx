import type { PropsWithChildren } from "react";
import { cn } from "~/lib/utils";

type Props = PropsWithChildren<{
  title: string;
  as?: "h1" | "h2";
}>;

const Heading = (props: Props) => {
  const { title, as = "h1" } = props;
  const Comp = as || "h1";
  return (
    <Comp
      className={cn("", {
        "scroll-m-20 text-4xl font-extrabold tracking-tight text-balance":
          as === "h1",
        "scroll-m-20 text-3xl font-semibold tracking-tight first:mt-0":
          as === "h2",
      })}
    >
      {title}
    </Comp>
  );
};

export default Heading;
