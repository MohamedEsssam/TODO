import * as Yup from "yup";

const todoSchema = Yup.object().shape({
  text: Yup.string().required("todo is required"),
});

export default todoSchema;
