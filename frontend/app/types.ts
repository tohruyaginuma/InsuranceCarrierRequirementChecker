export type Applicant = {
  id: number;
  given_name: string;
  surname: string;
  date_of_birth: string;
  insurance_status: string;
  prior_carrier: string;
  umpd: number | null;
  collision: number | null;
};
