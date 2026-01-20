import Heading from "~/components/heading";
import type { Route } from "./+types/create";
import Content from "~/components/content";
import Flex from "~/components/flex";
import { Link } from "react-router";
import { ArrowLeftIcon } from "lucide-react";
import { Input } from "~/components/ui/input";
import { Button } from "~/components/ui/button";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectLabel,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/components/ui/select";
import data from "../../../data/data.json";
import { apiUrl } from "~/lib/config";
import {
  Field,
  FieldContent,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSeparator,
  FieldSet,
  FieldTitle,
} from "~/components/ui/field";
import { toast } from "sonner";
import { useState } from "react";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Create | Insurance Carrier Requirement Checker" },
    { name: "description", content: "Create a new applicant" },
  ];
}

export default function Home() {
  const [errors, setErrors] = useState<Record<string, string>>({});

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const data = {
      given_name: formData.get("ApplicantGivenName"),
      surname: formData.get("ApplicantSurname"),
      date_of_birth: formData.get("ApplicantDOB"),
      insurance_status: formData.get("InsuranceStatus"),
      prior_carrier: formData.get("PriorCarrier"),
      umpd: formData.get("UMPD"),
      collision: formData.get("Collision"),
    };
    console.log(data);

    try {
      const response = await fetch(`${apiUrl}/v1/applicants/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      const result = await response.json();
      console.log(result);
      if (result.result === "OK") {
        toast.success("Applicant created successfully");
      } else {
        toast.error(result.message || "Failed to create applicant");
      }
    } catch (error) {
      toast.error(
        error instanceof Error ? error.message : "Failed to create applicant",
      );
    }
  };

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
      <Content>
        <FieldSet className="max-w-2xl">
          <FieldLegend>Profile</FieldLegend>
          <FieldDescription>
            This appears on invoices and emails.
          </FieldDescription>
          <form onSubmit={handleSubmit}>
            <FieldGroup>
              <div className="grid grid-cols-2 gap-x-6 gap-y-6">
                {data.fieldDefinitions.map((field) => (
                  <Field key={field.propertyName}>
                    <FieldLabel htmlFor={field.propertyName}>
                      {field.label}
                    </FieldLabel>

                    {field.options ? (
                      <Select>
                        <SelectTrigger className="w-full">
                          <SelectValue
                            placeholder={
                              data.examples[0].values
                                .find(
                                  (v) => v.propertyName === field.propertyName,
                                )
                                ?.value?.toString() || ""
                            }
                          />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            {field.options.map((option) => (
                              <SelectItem
                                key={option.value}
                                value={option.value.toString()}
                              >
                                {option.label}
                              </SelectItem>
                            ))}
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    ) : (
                      <>
                        {field.type === "string" && (
                          <Input
                            type="text"
                            placeholder={
                              data.examples[0].values
                                .find(
                                  (v) => v.propertyName === field.propertyName,
                                )
                                ?.value?.toString() || ""
                            }
                            autoComplete="off"
                          />
                        )}
                      </>
                    )}
                  </Field>
                ))}
              </div>
              <Field orientation="horizontal">
                <Button type="submit">Submit</Button>
                <Button variant="outline" type="button">
                  Cancel
                </Button>
              </Field>
            </FieldGroup>
          </form>
        </FieldSet>
      </Content>
    </>
  );
}
