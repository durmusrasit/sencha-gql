type Theme = {
  themeName: string;
  backgroundColor: string;
  foregroundColor: string;
};

type CreateThemeInput = Pick<
  Theme,
  "themeName" | "backgroundColor" | "foregroundColor"
>;

type UpdateThemeInput = Pick<
  Theme,
  "themeName" | "backgroundColor" | "foregroundColor"
>;

type DeleteThemeInput = Pick<Theme, "themeName">;

type GetThemeInput = Pick<Theme, "themeName">;

export type {
  Theme,
  CreateThemeInput,
  UpdateThemeInput,
  DeleteThemeInput,
  GetThemeInput,
};
