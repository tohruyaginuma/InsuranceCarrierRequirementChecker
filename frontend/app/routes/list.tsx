import Heading from "~/components/heading";
import type { Route } from "./+types/list";
import {
	Table,
	TableHeader,
	TableRow,
	TableHead,
	TableCaption,
	TableBody,
	TableCell,
} from "~/components/ui/table";
import Content from "~/components/content";
import Flex from "~/components/flex";
import { Button } from "~/components/ui/button";
import { Link } from "react-router";
import { PlusIcon } from "lucide-react";

export function meta({}: Route.MetaArgs) {
	return [
		{ title: "List | Insurance Carrier Requirement Checker" },
		{ name: "description", content: "List of applicants" },
	];
}

export default function List() {
	return (
		<>
			<Flex direction="row" gap="md" justify="between" align="center">
				<Heading title="Applicants List" />
				<Button variant="outline" asChild>
					<Link to="/create">
						<PlusIcon className="size-4" />
						Create Applicant
					</Link>
				</Button>
			</Flex>
			<Content>
				<Table>
					<TableCaption>A list of your recent applicants.</TableCaption>
					<TableHeader>
						<TableRow>
							<TableHead className="w-[100px]">Invoice</TableHead>
							<TableHead>Status</TableHead>
							<TableHead>Method</TableHead>
							<TableHead className="text-right">Amount</TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						<TableRow>
							<TableCell className="font-medium">INV001</TableCell>
							<TableCell>Paid</TableCell>
							<TableCell>Credit Card</TableCell>
							<TableCell className="text-right">$250.00</TableCell>
						</TableRow>
					</TableBody>
				</Table>
			</Content>
		</>
	);
}
