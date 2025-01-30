import { useForm } from "react-hook-form";
import { UseMutateFunction } from "@tanstack/react-query";
import { z, ZodSchema } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

const useZodform = (
  schema: ZodSchema,
  mutation: UseMutateFunction,
  defaultValues?: any
) => {
  type Schema = z.infer<typeof schema>;
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors },
  } = useForm<Schema>({
    resolver: zodResolver(schema),
    defaultValues: { ...defaultValues },
  });

  const onFormSubmit = handleSubmit(async (values) => mutation({ ...values }));

  return { register, onFormSubmit, errors, watch, reset };
};

export default useZodform;
