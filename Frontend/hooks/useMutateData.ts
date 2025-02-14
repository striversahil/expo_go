import { useMutation, useQueryClient } from "@tanstack/react-query";

const useMutateData = (url: string) => {
  const queryClient = useQueryClient();
  const { mutate } = useMutation(
    (data: any) =>
      fetch(url, { method: "POST", body: JSON.stringify(data) }).then((res) =>
        res.json()
      ),
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["data"]);
      },
    }
  );
  return mutate;
};
export default useMutateData;
