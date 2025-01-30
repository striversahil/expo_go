import { View, Text } from "react-native";
import React from "react";
import { Link } from "expo-router";
import LoginScreen from "../(auth)/_components/layout";

type Props = {};

const Page = (props: Props) => {
  return (
    <View>
      <View>
        <LoginScreen screen={"Signup"} />
      </View>
    </View>
  );
};

export default Page;
