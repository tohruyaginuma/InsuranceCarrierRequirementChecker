import type { PropsWithChildren } from "react";
import { cn } from "~/lib/utils";

type Props = PropsWithChildren<{
	direction?: "row" | "column";
	justify?: "start" | "end" | "center" | "between" | "around" | "evenly";
	align?: "start" | "end" | "center" | "baseline" | "stretch";
	gap?: "sm" | "md" | "lg";
	className?: string;
}>;

const Flex = (props: Props) => {
	const {
		children,
		direction = "row",
		justify = "start",
		align = "start",
		gap = "0",
		className = "",
	} = props;
	return (
		<div
			className={cn("flex", {
				"flex-row": direction === "row",
				"flex-col": direction === "column",
				"justify-start": justify === "start",
				"justify-end": justify === "end",
				"justify-center": justify === "center",
				"justify-between": justify === "between",
				"justify-around": justify === "around",
				"justify-evenly": justify === "evenly",
				"items-start": align === "start",
				"items-end": align === "end",
				"items-center": align === "center",
				"items-baseline": align === "baseline",
				"items-stretch": align === "stretch",
				"gap-sm": gap === "sm",
				"gap-md": gap === "md",
				"gap-lg": gap === "lg",
				className,
			})}
		>
			{children}
		</div>
	);
};

export default Flex;
