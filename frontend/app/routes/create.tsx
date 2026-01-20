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
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/components/ui/select";
import data from "../../../data/data.json";
import { apiUrl } from "~/lib/config";
import {
  Field,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSet,
} from "~/components/ui/field";
import { toast } from "sonner";
import { useState } from "react";
import type { FieldValidationResult } from "~/types";
import { useNavigate } from "react-router";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Create | Insurance Carrier Requirement Checker" },
    { name: "description", content: "Create a new applicant" },
  ];
}

export default function Home() {
  const navigate = useNavigate();

  const [errors, setErrors] = useState<FieldValidationResult[]>([]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const data = {
      given_name: formData.get("ApplicantGivenName"),
      surname: formData.get("ApplicantSurname"),
      date_of_birth: formData.get("ApplicantDOB"),
      insurance_status: formData.get("InsuranceStatus"),
      prior_carrier: formData.get("PriorCarrier"),
      umpd: formData.get("UMPD")
        ? parseInt(formData.get("UMPD") as string)
        : null,
      collision: formData.get("Collision")
        ? parseInt(formData.get("Collision") as string)
        : null,
    };

    try {
      const response = await fetch(`${apiUrl}/v1/applicants/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      const result = await response.json();
      setErrors([]);

      if (result.result === "NG") {
        if (result.fieldResults) {
          setErrors(result.fieldResults);
        }

        toast.error(result.message || "Failed to create applicant");
        return;
      } else {
        toast.success("Applicant created successfully");
        navigate("/");
      }
    } catch (error) {
      setErrors([]);
      toast.error(
        error instanceof Error ? error.message : "Failed to create applicant",
      );
    }
  };

  const handleCancel = () => {
    navigate("/");
  };

  return (
    <>
      <Link to="/" className="block mb-4 hover:opacity-60 transition-opacity duration-100">
        <Heading title="Insurance Carrier Requirement Checker" as="h1" />
      </Link>
      <Flex direction="row" gap="md" justify="between" align="center">
        <Heading title="Create Applicant" as="h2" />
        <Button variant="outline" asChild>
          <Link to="/">
            <ArrowLeftIcon className="size-4" />
            Back to List
          </Link>
        </Button>
      </Flex>
      <Content>
        <FieldSet className="w-1/2">
          <FieldLegend>Applicant</FieldLegend>
          <FieldDescription>
            This is the applicant's information.
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
                      <Select name={field.propertyName}>
                        <SelectTrigger className="w-full">
                          <SelectValue placeholder="Select an option" />
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
                            type={
                              field.propertyName === "ApplicantDOB"
                                ? "date"
                                : "text"
                            }
                            name={field.propertyName}
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
                    {(() => {
                      const error = errors?.find(
                        (error) => error.propertyName === field.propertyName,
                      );

                      return (
                        error && <FieldError>{error.errorMessage}</FieldError>
                      );
                    })()}
                  </Field>
                ))}
              </div>
              <Field orientation="horizontal">
                <Button type="submit">Submit</Button>
                <Button variant="outline" type="button" onClick={handleCancel}>
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
