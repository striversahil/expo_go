// Explore screen with collapsible sections
import { StyleSheet, Image, Platform, View, Text } from "react-native";

import { Collapsible } from "@/components/Collapsible";
import { ExternalLink } from "@/components/ExternalLink";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { IconSymbol } from "@/components/ui/IconSymbol";
import { HelloWave } from "@/components/HelloWave";

export default function TabTwoScreen() {
  return (
    <View style={{ flex: 1 }}>
      <Text className="text-2xl text-red-600 bg-green-400 rounded-lg p-4 text-center font-bold">
        Explore
      </Text>
    </View>
  );
}

const styles = StyleSheet.create({
  headerImage: {
    color: "#808080",
    bottom: -90,
    left: -35,
    position: "absolute",
  },
  titleContainer: {
    flexDirection: "row",
    gap: 8,
  },
});
