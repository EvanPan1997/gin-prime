import {ThemeConfig} from "antd/es/config-provider/context";
import {defaultTheme} from "antd/lib/theme/context";
import {theme} from "antd";


const colorPrimaryConfig = {
    colorPrimaryBg: "#e6f4ff",
    colorPrimaryBgHover: "#bae0ff",
    colorPrimaryBorder: "#91caff",
    colorPrimaryBorderHover: "#69b1ff",
    colorPrimaryHover: "#4096ff",
    colorPrimary: "#1677FF",
    colorPrimaryActive: "#0958d9",
    colorPrimaryTextHover: "#4096ff",
    colorPrimaryText: "#1677ff",
    colorPrimaryTextActive: "#0958d9",
}

const colorSuccessConfig = {
    colorSuccessBg: "#f6ffed",
    colorSuccessBgHover: "#d9f7be",
    colorSuccessBorder: "#b7eb8f",
    colorSuccessBorderHover: "#95de64",
    colorSuccessHover: "#95de64",
    colorSuccess: "#52c41a",
    colorSuccessActive: "#389e0d",
    colorSuccessTextHover: "#73d13d",
    colorSuccessText: "#52c41a",
    colorSuccessTextActive: "#389e0d",
}

const colorWarningConfig = {
    colorWarningBg: "#fffbe6",
    colorWarningBgHover: "#fff1b8",
    colorWarningBorder: "#ffe58f",
    colorWarningBorderHover: "#ffd666",
    colorWarningHover: "#ffd666",
    colorWarning: "#faad14",
    colorWarningActive: "#d48806",
    colorWarningTextHover: "#ffc53d",
    colorWarningText: "#faad14",
    colorWarningTextActive: "#d48806",
}

const colorErrorConfig = {
    colorErrorBg: "#fff2f0",
    colorErrorBgHover: "#fff1f0",
    colorErrorBorder: "#ffccc7",
    colorErrorBorderHover: "#ffa39e",
    colorErrorHover: "#ff7875",
    colorError: "#ff4d4f",
    colorErrorActive: "#d9363e",
    colorErrorTextHover: "#ff7875",
    colorErrorText: "#ff4d4f",
    colorErrorTextActive: "#d9363e",
}

const colorLinkConfig = {
    colorLink: "#1677ff",
    colorLinkHover: "#69b1ff",
    colorLinkActive: "#0958d9"
}

// 中性色
const colorBaseConfig = {
    colorTextBase: "#000000",
    colorBgBase: "#ffffff",
    // 文本
    colorText: "rgba(0, 0, 0, 0.88)",
    colorTextSecondary: "rgba(0, 0, 0, 0.65)",
    colorTextTertiary: "rgba(0, 0, 0, 0.45)",
    colorTextQuaternary: "rgba(0, 0, 0, 0.25)",
    // 描边
    colorBorder: "#d9d9d9",
    colorBorderSecondary: "#f0f0f0",
    // 填充
    colorFill: "rgba(0, 0, 0, 0.15)",
    colorFillSecondary: "rgba(0, 0, 0, 0.06)",
    colorFillTertiary: "rgba(0, 0, 0, 0.04)",
    colorFillQuaternary: "rgba(0, 0, 0, 0.02)",
    // 背景
    colorBgContainer: "#ffffff",
    colorBgElevated: "#ffffff",
    colorBgLayout: "#f5f5f5",
    colorBgSpotlight: "rgba(0, 0, 0, 0.85)",
    colorBgMask: "rgba(0, 0, 0, 0.45)",
}

const sizeConfig = {
    fontSize: 14,
    fontSizeSM: 12,
    fontSizeLG: 16,
    fontSizeXL: 20,

    fontSizeHeading1: 38,
    fontSizeHeading2: 30,
    fontSizeHeading3: 24,
    fontSizeHeading4: 20,
    fontSizeHeading5: 16,

    lineHeight: 1.5714285714285714,
    lineHeightSM: 1.6666666666666667,
    lineHeightLG: 1.5,

    lineHeightHeading1: 1.2105263157894737,
    lineHeightHeading2: 1.2666666666666666,
    lineHeightHeading3: 1.3333333333333333,
    lineHeightHeading4: 1.4,
    lineHeightHeading5: 1.5,

    sizeStep: 4,
    sizeUnit: 4,
}

const borderRadiusConfig = {
    borderRadiusXS: 2,
    borderRadiusSM: 4,
    borderRadius: 6,
    borderRadiusLG: 8,
}

const boxShadowConfig = {
    boxShadow: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)",
    boxShadowSecondary: "0 6px 16px 0 rgba(0, 0, 0, 0.08), 0 3px 6px -4px rgba(0, 0, 0, 0.12), 0 9px 28px 8px rgba(0, 0, 0, 0.05)",
}

const lightThemeConfig: ThemeConfig = {
    ...defaultTheme,
    token: {
        ...colorPrimaryConfig,
        ...colorSuccessConfig,
        ...colorWarningConfig,
        ...colorErrorConfig,
        ...colorLinkConfig,
        ...colorBaseConfig,
        ...sizeConfig,
        ...borderRadiusConfig,
        ...boxShadowConfig,
    },
    components: {
        Layout: {
            headerBg: "#f5f5f5",
            // bodyBg: '#ffffff',
            // headerPadding: 0,
            // siderBg: "#d9d9d9",
        },
    },
}

const darkThemeConfig: ThemeConfig = {
    ...defaultTheme,
    algorithm: theme.darkAlgorithm,
    components: {
        Layout: {
            headerBg: "#000000",
        }
    }
}

export function CustomTheme(isDark: boolean) {
    return isDark ? darkThemeConfig : lightThemeConfig
}

export function MainContentBorder(isDark: boolean): any {
    const cssProps = {
        margin: "16px 16px 0px 16px",
        padding: '8px',
        borderRadius: "8px",
        overflowY: "auto",
        // border: '0.2px solid gray'
    }
    return isDark ? {...cssProps, background: '#141414'}
        :{...cssProps, background: '#ffffff'}
}