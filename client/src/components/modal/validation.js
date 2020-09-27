import * as Yup from "yup";

const TodoSchema = Yup.object().shape({
  name: Yup.string().required(),
});

export default TodoSchema;
