import * as React from "react";
import OrderCard, { OrderItem } from "./OrderCard";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/PaymentPage/OrderCard",
  decorators: [withKnobs],
};

const defaultOrderItems: OrderItem[] = [
    {
        id: '082739428373', name: 'Cargo Manager (CM) – VC, HS, XRY, EDS', quantity: 1, price: 200, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    },
    {
        id: '082739428374', name: 'Cargo Aircraft Protection', quantity: 3, price: 55, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    },
    {
        id: '082739428375', name: 'Cargo Manager Recurrent (CM) – VC, HS, XRY, EDS', quantity: 2, price: 25, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")
    }
];

export const normal = () => {
    return <OrderCard orderItems={defaultOrderItems} />;
  };