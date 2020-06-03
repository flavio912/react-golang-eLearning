import React, { useState, ChangeEvent } from "react";
import { createUseStyles } from "react-jss";
import classNames from "classnames";

const styles = createUseStyles({
  defaultStyles: {
    border: "none",
    outline: "none",
    flex: 1,
  },
});

export type InputTypes =
  | "text"
  | "password"
  | "email"
  | "search"
  | "tel"
  | "number";
type Props = {
  type: InputTypes;
  className?: string;
  placeholder?: string;
  onChange: (text: string) => string | Promise<void> | void; // Function given changed text and should return what to update it to
  onFocus?: () => void;
  onBlur?: () => void;
  value?: string;
  setValue?: (text: string) => any;
};

const CoreInput = React.forwardRef(
  (
    {
      type,
      className,
      placeholder = "",
      onChange,
      onFocus,
      onBlur,
      value,
      setValue,
    }: Props,
    ref: React.Ref<HTMLInputElement>
  ) => {
    const classes = styles();

    let _value, _setValue: (text: string) => any;
    if (value !== undefined && setValue) {
      _value = value;
      _setValue = setValue;
    } else {
      [_value, _setValue] = useState("");
    }

    const updateInput = (event: ChangeEvent<HTMLInputElement>) => {
      let newValue = onChange(event.target.value) || event.target.value;

      // Only update with given value if it is a string
      if (typeof newValue != "string") {
        newValue = event.target.value;
      }
      _setValue(newValue);
    };

    return (
      <input
        type={type}
        className={classNames(className, classes.defaultStyles)}
        value={_value}
        placeholder={placeholder}
        onChange={updateInput}
        ref={ref}
        onFocus={onFocus}
        onBlur={onBlur}
      />
    );
  }
);

export default CoreInput;
