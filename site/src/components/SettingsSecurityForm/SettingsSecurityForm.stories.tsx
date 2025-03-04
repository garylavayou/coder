import { Story } from "@storybook/react"
import { SecurityForm, SecurityFormProps } from "./SettingsSecurityForm"
import { mockApiError } from "testHelpers/entities"

export default {
  title: "components/SettingsSecurityForm",
  component: SecurityForm,
  argTypes: {
    onSubmit: { action: "Submit" },
  },
}

const Template: Story<SecurityFormProps> = (args: SecurityFormProps) => (
  <SecurityForm {...args} />
)

export const Example = Template.bind({})
Example.args = {
  isLoading: false,
  initialValues: {
    old_password: "",
    password: "",
    confirm_password: "",
  },
  updateSecurityError: undefined,
  onSubmit: () => {
    return Promise.resolve()
  },
}

export const Loading = Template.bind({})
Loading.args = {
  ...Example.args,
  isLoading: true,
}

export const WithError = Template.bind({})
WithError.args = {
  ...Example.args,
  updateSecurityError: mockApiError({
    message: "Old password is incorrect",
    validations: [
      {
        field: "old_password",
        detail: "Old password is incorrect.",
      },
    ],
  }),
  initialTouched: {
    old_password: true,
  },
}
