import React from "react";
import { Field, useFormikContext } from "formik";
import "./style.css";

const CustomInput = ({ initialValues, name, ...otherProps }) => {
  const { handleChange, setFieldTouched, touched, errors } = useFormikContext();

  return (
    <>
      <Field
        defaultValue={initialValues}
        id="todoInput"
        className="form-control"
        name={name}
        onBlur={() => setFieldTouched(name)}
        onChange={handleChange(name)}
        {...otherProps}
      />
    </>
  );
};

export default CustomInput;
