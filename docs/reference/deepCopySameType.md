# Setting: deepCopySameType

`deepCopySameType [yes,no]` is a [boolean setting](./define-settings.md#boolean)
and can be defined as [CLI argument](./define-settings.md#cli), [conversion
comment](./define-settings.md#conversion) or [method
comment](./define-settings.md#method). This setting is
[inheritable](./define-settings.md#inheritance). Default `true`.

Goverter deep copies instances by default when converting the source to the
target type. You can disable deep copying with `goverter:deepCopySameType no`.

::: code-group
<<< @../../example/skip-deep-copy-same-type/input.go
<<< @../../example/skip-deep-copy-same-type/generated/generated.go [generated/generated.go]
:::
