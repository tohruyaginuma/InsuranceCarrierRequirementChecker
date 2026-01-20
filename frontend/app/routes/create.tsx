import Heading from "~/components/heading";
import type { Route } from "./+types/create";
import Content from "~/components/content";
import Flex from "~/components/flex";
import { Button } from "~/components/ui/button";
import { Link } from "react-router";
import { ArrowLeftIcon } from "lucide-react";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Create | Insurance Carrier Requirement Checker" },
    { name: "description", content: "Create a new applicant" },
  ];
}

export default function Home() {
  return (
    <>
      <Flex direction="row" gap="md" justify="between" align="center">
        <Heading title="Create Applicant" />
        <Button variant="outline" asChild>
          <Link to="/">
            <ArrowLeftIcon className="size-4" />
            Back to List
          </Link>
        </Button>
      </Flex>
      <Content>content</Content>
    </>
  );
}
