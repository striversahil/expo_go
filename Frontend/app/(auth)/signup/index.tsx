import { StyleSheet, Text, View } from "react-native";
import React from "react";
import LoginScreen from "../_components/layout";

type Props = {};

const Page = ({ navigation }: { navigation: any }) => {
  return (
    <View className="h-screen">
      <Text onPress={() => navigation.navigate("Login")} className="text-white">
        Click Here
      </Text>
      <LoginScreen screen={"Signup"} />
    </View>
  );
};

export default Page;

const styles = StyleSheet.create({});
