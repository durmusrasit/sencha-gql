import { themes } from "../db";
import {
  CreateThemeInput,
  UpdateThemeInput,
  DeleteThemeInput,
  GetThemeInput,
  Theme,
} from "../types";

const createTheme = (args: { input: CreateThemeInput }): Theme => {
  const theme = args.input;

  themes.push(theme);

  return theme;
};

const updateTheme = (args: { input: UpdateThemeInput }): boolean => {
  return true;
};

const deleteTheme = (args: { input: DeleteThemeInput }): boolean => {
  return true;
};

const getTheme = (args: { input: GetThemeInput }): Theme | undefined =>
  themes.find((t: Theme) => t.themeName === args.input.themeName);

const getThemes = (): Theme[] => themes;

const root = {
  createTheme,
  updateTheme,
  deleteTheme,
  getTheme,
  getThemes,
};

export { root };
