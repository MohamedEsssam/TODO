import React from "react";
import { Field, ErrorMessage, useFormikContext } from "formik";

const CustomInput = ({ initialValues, name, ...otherProps }) => {
  const { handleChange, setFieldTouched, touched, errors } = useFormikContext();

  return (
    <>
      <Field
        defaultValue={initialValues}
        id="input"
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
