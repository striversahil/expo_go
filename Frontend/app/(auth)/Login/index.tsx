import { StyleSheet, Text, View } from "react-native";
import React from "react";
import LoginScreen from "../_components/layout";

type Props = {};

const Page = (props: Props) => {
  return (
    <View className="">
      <LoginScreen scrreen={"Login"} />
    </View>
  );
};

export default Page;

const styles = StyleSheet.create({});
