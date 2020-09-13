import React from "react";
import { Field, ErrorMessage, useFormikContext } from "formik";

const CustomInput = ({ initialValues, name, ...otherProps }) => {
  const { handleChange, setFieldTouched, touched, errors } = useFormikContext();

  return (
    <>
      <label id="label">{name} :</label>
      <Field
        defaultValue={initialValues}
        id="input"
        className="form-control"
        name={name}
        onBlur={() => setFieldTouched(name)}
        onChange={handleChange(name)}
        {...otherProps}
      />
      {touched[name] && <ErrorMessage id="error" name={name} component="div" />}
    </>
  );
};

export default CustomInput;
