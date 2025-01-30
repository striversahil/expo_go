// Explore screen with collapsible sections

import { Text, View, Button, StyleSheet } from "react-native";

import { Collapsible } from "@/components/Collapsible";
import { ExternalLink } from "@/components/ExternalLink";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { IconSymbol } from "@/components/ui/IconSymbol";
import { HelloWave } from "@/components/HelloWave";

import { useForm, Controller, SubmitHandler } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { TextInput } from "react-native-paper";

export default function TabsScreen() {
  const userSchema = z.object({
    username: z
      .string()
      .min(5, { message: "Username must be at least 5 characters." })
      .max(30, { message: "Username must be at most 30 characters." }),
    email: z.string().email({ message: "Please enter a valid email address." }),
    age: z.number().min(18, { message: "You must be over 18 years old." }),
  });

  type UserFormType = z.infer<typeof userSchema>;

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<UserFormType>({
    resolver: zodResolver(userSchema),
  });

  const onSubmit: SubmitHandler<UserFormType> = (data: UserFormType) => {
    console.log(data);
  };

  return (
    <View style={styles.container} className="bg-white">
      <Controller
        control={control}
        render={({
          field: { onChange, onBlur, value },
          fieldState: { error },
        }) => (
          <View>
            <TextInput
              mode="outlined"
              label="Username"
              onBlur={onBlur}
              onChangeText={onChange}
              value={value}
              activeOutlineColor={error ? "red" : "black"}
              style={styles.input}
            />
          </View>
        )}
        name="username"
      />
      {errors.username && (
        <Text style={{ color: "#ff8566" }}>{errors.username.message}</Text>
      )}

      <Controller
        control={control}
        render={({
          field: { onChange, onBlur, value },
          fieldState: { error },
        }) => (
          <TextInput
            mode="outlined"
            label="E-mail"
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            activeOutlineColor={error ? "red" : "black"}
            style={styles.input}
          />
        )}
        name="email"
      />
      {errors.email && (
        <Text style={{ color: "#ff8566" }}>{errors.email.message}</Text>
      )}

      <Controller
        control={control}
        render={({
          field: { onChange, onBlur, value },
          fieldState: { error },
        }) => (
          <TextInput
            mode="outlined"
            label="Age"
            onBlur={onBlur}
            onChangeText={(text) => {
              const parsed = parseInt(text, 10);
              onChange(isNaN(parsed) ? "" : parsed);
            }}
            value={
              value === null || value === undefined ? "" : value.toString()
            }
            placeholder="Yaş"
            keyboardType="numeric"
            activeOutlineColor={error ? "red" : "black"}
            style={styles.input}
          />
        )}
        name="age"
      />
      {errors.age && (
        <Text style={{ color: "#ff8566" }}>{errors.age.message}</Text>
      )}

      <Button title="Submit" onPress={handleSubmit(onSubmit)} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    marginVertical: 10,
    padding: 8,
  },
  input: {
    marginVertical: 10,
  },
});
