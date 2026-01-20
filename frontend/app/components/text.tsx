import type { PropsWithChildren } from "react";
import { cn } from "~/lib/utils";

type Props = PropsWithChildren<{
  as?: "p" | "span";
  className?: string;
  isError?: boolean;
}>;
const Text = (props: Props) => {
  const { children, as = "p", isError = false } = props;
  const Comp = as || "p";

  return (
    <Comp className={cn("text-sm text-gray-500", { "text-red-500": isError })}>
      {children}
    </Comp>
  );
};

export default Text;
