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
import { useQuery } from "@tanstack/react-query";
import { apiUrl } from "~/lib/config";
import type { Applicant } from "~/types";
import { Spinner } from "~/components/ui/spinner";
import Text from "~/components/text";
import { useCallback } from "react";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "List | Insurance Carrier Requirement Checker" },
    { name: "description", content: "List of applicants" },
  ];
}

export default function List() {
  const fetchApplicants = useCallback(async () => {
    try {
      const response = await fetch(`${apiUrl}/v1/applicants/`);

      const data = await response.json();
      return data.applicants;
    } catch (error) {
      console.error(error);
      throw error;
    }
  }, []);

  const {
    data: applicants,
    isPending,
  } = useQuery<Applicant[]>({
    queryKey: ["applicants"],
    queryFn: fetchApplicants,
  });

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
              <TableHead className="w-[100px]">Given Name</TableHead>
              <TableHead>Surname</TableHead>
              <TableHead>Date of Birth</TableHead>
              <TableHead>Insurance Status</TableHead>
              <TableHead>Prior Carrier</TableHead>
              <TableHead>UMPD</TableHead>
              <TableHead>Collision</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {isPending ? (
              <TableRow key="loading">
                <TableCell colSpan={7} className="text-center">
                  <Flex justify="center" align="center" className="h-full">
                    <Spinner />
                  </Flex>
                </TableCell>
              </TableRow>
            ) : applicants && applicants.length > 0 ? (
              applicants.map((applicant) => (
                <TableRow key={applicant.id}>
                  <TableCell className="font-medium">
                    {applicant.given_name}
                  </TableCell>
                  <TableCell>{applicant.surname}</TableCell>
                  <TableCell>{applicant.date_of_birth}</TableCell>
                  <TableCell>{applicant.insurance_status}</TableCell>
                  <TableCell>{applicant.prior_carrier}</TableCell>
                  <TableCell>
                    {applicant.umpd ? applicant.umpd : "N/A"}
                  </TableCell>
                  <TableCell>
                    {applicant.collision ? applicant.collision : "N/A"}
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow key="no-applicants">
                <TableCell colSpan={7} className="text-center">
                  <Text as="span" isError>
                    No applicants found
                  </Text>
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </Content>
    </>
  );
}
