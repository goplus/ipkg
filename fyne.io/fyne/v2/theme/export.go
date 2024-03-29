// export by github.com/goplus/igop/cmd/qexp

package theme

import (
	q "fyne.io/fyne/v2/theme"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "theme",
		Path: "fyne.io/fyne/v2/theme",
		Deps: map[string]string{
			"encoding/hex":                 "hex",
			"encoding/json":                "json",
			"errors":                       "errors",
			"fyne.io/fyne/v2":              "fyne",
			"fyne.io/fyne/v2/internal/svg": "svg",
			"fyne.io/fyne/v2/storage":      "storage",
			"image/color":                  "color",
			"io":                           "io",
			"net/url":                      "url",
			"os":                           "os",
			"strings":                      "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"DisabledResource":       reflect.TypeOf((*q.DisabledResource)(nil)).Elem(),
			"ErrorThemedResource":    reflect.TypeOf((*q.ErrorThemedResource)(nil)).Elem(),
			"InvertedThemedResource": reflect.TypeOf((*q.InvertedThemedResource)(nil)).Elem(),
			"PrimaryThemedResource":  reflect.TypeOf((*q.PrimaryThemedResource)(nil)).Elem(),
			"ThemedResource":         reflect.TypeOf((*q.ThemedResource)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AccountIcon":               reflect.ValueOf(q.AccountIcon),
			"BackgroundColor":           reflect.ValueOf(q.BackgroundColor),
			"ButtonColor":               reflect.ValueOf(q.ButtonColor),
			"CancelIcon":                reflect.ValueOf(q.CancelIcon),
			"CaptionTextSize":           reflect.ValueOf(q.CaptionTextSize),
			"CheckButtonCheckedIcon":    reflect.ValueOf(q.CheckButtonCheckedIcon),
			"CheckButtonIcon":           reflect.ValueOf(q.CheckButtonIcon),
			"ColorAchromaticIcon":       reflect.ValueOf(q.ColorAchromaticIcon),
			"ColorChromaticIcon":        reflect.ValueOf(q.ColorChromaticIcon),
			"ColorPaletteIcon":          reflect.ValueOf(q.ColorPaletteIcon),
			"ComputerIcon":              reflect.ValueOf(q.ComputerIcon),
			"ConfirmIcon":               reflect.ValueOf(q.ConfirmIcon),
			"ContentAddIcon":            reflect.ValueOf(q.ContentAddIcon),
			"ContentClearIcon":          reflect.ValueOf(q.ContentClearIcon),
			"ContentCopyIcon":           reflect.ValueOf(q.ContentCopyIcon),
			"ContentCutIcon":            reflect.ValueOf(q.ContentCutIcon),
			"ContentPasteIcon":          reflect.ValueOf(q.ContentPasteIcon),
			"ContentRedoIcon":           reflect.ValueOf(q.ContentRedoIcon),
			"ContentRemoveIcon":         reflect.ValueOf(q.ContentRemoveIcon),
			"ContentUndoIcon":           reflect.ValueOf(q.ContentUndoIcon),
			"DarkTheme":                 reflect.ValueOf(q.DarkTheme),
			"DefaultSymbolFont":         reflect.ValueOf(q.DefaultSymbolFont),
			"DefaultTextBoldFont":       reflect.ValueOf(q.DefaultTextBoldFont),
			"DefaultTextBoldItalicFont": reflect.ValueOf(q.DefaultTextBoldItalicFont),
			"DefaultTextFont":           reflect.ValueOf(q.DefaultTextFont),
			"DefaultTextItalicFont":     reflect.ValueOf(q.DefaultTextItalicFont),
			"DefaultTextMonospaceFont":  reflect.ValueOf(q.DefaultTextMonospaceFont),
			"DefaultTheme":              reflect.ValueOf(q.DefaultTheme),
			"DeleteIcon":                reflect.ValueOf(q.DeleteIcon),
			"DisabledButtonColor":       reflect.ValueOf(q.DisabledButtonColor),
			"DisabledColor":             reflect.ValueOf(q.DisabledColor),
			"DisabledTextColor":         reflect.ValueOf(q.DisabledTextColor),
			"DocumentCreateIcon":        reflect.ValueOf(q.DocumentCreateIcon),
			"DocumentIcon":              reflect.ValueOf(q.DocumentIcon),
			"DocumentPrintIcon":         reflect.ValueOf(q.DocumentPrintIcon),
			"DocumentSaveIcon":          reflect.ValueOf(q.DocumentSaveIcon),
			"DownloadIcon":              reflect.ValueOf(q.DownloadIcon),
			"ErrorColor":                reflect.ValueOf(q.ErrorColor),
			"ErrorIcon":                 reflect.ValueOf(q.ErrorIcon),
			"FileApplicationIcon":       reflect.ValueOf(q.FileApplicationIcon),
			"FileAudioIcon":             reflect.ValueOf(q.FileAudioIcon),
			"FileIcon":                  reflect.ValueOf(q.FileIcon),
			"FileImageIcon":             reflect.ValueOf(q.FileImageIcon),
			"FileTextIcon":              reflect.ValueOf(q.FileTextIcon),
			"FileVideoIcon":             reflect.ValueOf(q.FileVideoIcon),
			"FocusColor":                reflect.ValueOf(q.FocusColor),
			"FolderIcon":                reflect.ValueOf(q.FolderIcon),
			"FolderNewIcon":             reflect.ValueOf(q.FolderNewIcon),
			"FolderOpenIcon":            reflect.ValueOf(q.FolderOpenIcon),
			"ForegroundColor":           reflect.ValueOf(q.ForegroundColor),
			"FromJSON":                  reflect.ValueOf(q.FromJSON),
			"FromJSONReader":            reflect.ValueOf(q.FromJSONReader),
			"FromLegacy":                reflect.ValueOf(q.FromLegacy),
			"FyneLogo":                  reflect.ValueOf(q.FyneLogo),
			"GridIcon":                  reflect.ValueOf(q.GridIcon),
			"HelpIcon":                  reflect.ValueOf(q.HelpIcon),
			"HistoryIcon":               reflect.ValueOf(q.HistoryIcon),
			"HomeIcon":                  reflect.ValueOf(q.HomeIcon),
			"HoverColor":                reflect.ValueOf(q.HoverColor),
			"IconInlineSize":            reflect.ValueOf(q.IconInlineSize),
			"InfoIcon":                  reflect.ValueOf(q.InfoIcon),
			"InnerPadding":              reflect.ValueOf(q.InnerPadding),
			"InputBackgroundColor":      reflect.ValueOf(q.InputBackgroundColor),
			"InputBorderColor":          reflect.ValueOf(q.InputBorderColor),
			"InputBorderSize":           reflect.ValueOf(q.InputBorderSize),
			"LightTheme":                reflect.ValueOf(q.LightTheme),
			"LineSpacing":               reflect.ValueOf(q.LineSpacing),
			"ListIcon":                  reflect.ValueOf(q.ListIcon),
			"LoginIcon":                 reflect.ValueOf(q.LoginIcon),
			"LogoutIcon":                reflect.ValueOf(q.LogoutIcon),
			"MailAttachmentIcon":        reflect.ValueOf(q.MailAttachmentIcon),
			"MailComposeIcon":           reflect.ValueOf(q.MailComposeIcon),
			"MailForwardIcon":           reflect.ValueOf(q.MailForwardIcon),
			"MailReplyAllIcon":          reflect.ValueOf(q.MailReplyAllIcon),
			"MailReplyIcon":             reflect.ValueOf(q.MailReplyIcon),
			"MailSendIcon":              reflect.ValueOf(q.MailSendIcon),
			"MediaFastForwardIcon":      reflect.ValueOf(q.MediaFastForwardIcon),
			"MediaFastRewindIcon":       reflect.ValueOf(q.MediaFastRewindIcon),
			"MediaMusicIcon":            reflect.ValueOf(q.MediaMusicIcon),
			"MediaPauseIcon":            reflect.ValueOf(q.MediaPauseIcon),
			"MediaPhotoIcon":            reflect.ValueOf(q.MediaPhotoIcon),
			"MediaPlayIcon":             reflect.ValueOf(q.MediaPlayIcon),
			"MediaRecordIcon":           reflect.ValueOf(q.MediaRecordIcon),
			"MediaReplayIcon":           reflect.ValueOf(q.MediaReplayIcon),
			"MediaSkipNextIcon":         reflect.ValueOf(q.MediaSkipNextIcon),
			"MediaSkipPreviousIcon":     reflect.ValueOf(q.MediaSkipPreviousIcon),
			"MediaStopIcon":             reflect.ValueOf(q.MediaStopIcon),
			"MediaVideoIcon":            reflect.ValueOf(q.MediaVideoIcon),
			"MenuBackgroundColor":       reflect.ValueOf(q.MenuBackgroundColor),
			"MenuDropDownIcon":          reflect.ValueOf(q.MenuDropDownIcon),
			"MenuDropUpIcon":            reflect.ValueOf(q.MenuDropUpIcon),
			"MenuExpandIcon":            reflect.ValueOf(q.MenuExpandIcon),
			"MenuIcon":                  reflect.ValueOf(q.MenuIcon),
			"MoreHorizontalIcon":        reflect.ValueOf(q.MoreHorizontalIcon),
			"MoreVerticalIcon":          reflect.ValueOf(q.MoreVerticalIcon),
			"MoveDownIcon":              reflect.ValueOf(q.MoveDownIcon),
			"MoveUpIcon":                reflect.ValueOf(q.MoveUpIcon),
			"NavigateBackIcon":          reflect.ValueOf(q.NavigateBackIcon),
			"NavigateNextIcon":          reflect.ValueOf(q.NavigateNextIcon),
			"NewDisabledResource":       reflect.ValueOf(q.NewDisabledResource),
			"NewErrorThemedResource":    reflect.ValueOf(q.NewErrorThemedResource),
			"NewInvertedThemedResource": reflect.ValueOf(q.NewInvertedThemedResource),
			"NewPrimaryThemedResource":  reflect.ValueOf(q.NewPrimaryThemedResource),
			"NewThemedResource":         reflect.ValueOf(q.NewThemedResource),
			"OverlayBackgroundColor":    reflect.ValueOf(q.OverlayBackgroundColor),
			"Padding":                   reflect.ValueOf(q.Padding),
			"PlaceHolderColor":          reflect.ValueOf(q.PlaceHolderColor),
			"PressedColor":              reflect.ValueOf(q.PressedColor),
			"PrimaryColor":              reflect.ValueOf(q.PrimaryColor),
			"PrimaryColorNamed":         reflect.ValueOf(q.PrimaryColorNamed),
			"PrimaryColorNames":         reflect.ValueOf(q.PrimaryColorNames),
			"QuestionIcon":              reflect.ValueOf(q.QuestionIcon),
			"RadioButtonCheckedIcon":    reflect.ValueOf(q.RadioButtonCheckedIcon),
			"RadioButtonIcon":           reflect.ValueOf(q.RadioButtonIcon),
			"ScrollBarColor":            reflect.ValueOf(q.ScrollBarColor),
			"ScrollBarSize":             reflect.ValueOf(q.ScrollBarSize),
			"ScrollBarSmallSize":        reflect.ValueOf(q.ScrollBarSmallSize),
			"SearchIcon":                reflect.ValueOf(q.SearchIcon),
			"SearchReplaceIcon":         reflect.ValueOf(q.SearchReplaceIcon),
			"SelectionColor":            reflect.ValueOf(q.SelectionColor),
			"SeparatorColor":            reflect.ValueOf(q.SeparatorColor),
			"SeparatorThicknessSize":    reflect.ValueOf(q.SeparatorThicknessSize),
			"SettingsIcon":              reflect.ValueOf(q.SettingsIcon),
			"ShadowColor":               reflect.ValueOf(q.ShadowColor),
			"StorageIcon":               reflect.ValueOf(q.StorageIcon),
			"SuccessColor":              reflect.ValueOf(q.SuccessColor),
			"TextBoldFont":              reflect.ValueOf(q.TextBoldFont),
			"TextBoldItalicFont":        reflect.ValueOf(q.TextBoldItalicFont),
			"TextColor":                 reflect.ValueOf(q.TextColor),
			"TextFont":                  reflect.ValueOf(q.TextFont),
			"TextHeadingSize":           reflect.ValueOf(q.TextHeadingSize),
			"TextItalicFont":            reflect.ValueOf(q.TextItalicFont),
			"TextMonospaceFont":         reflect.ValueOf(q.TextMonospaceFont),
			"TextSize":                  reflect.ValueOf(q.TextSize),
			"TextSubHeadingSize":        reflect.ValueOf(q.TextSubHeadingSize),
			"UploadIcon":                reflect.ValueOf(q.UploadIcon),
			"ViewFullScreenIcon":        reflect.ValueOf(q.ViewFullScreenIcon),
			"ViewRefreshIcon":           reflect.ValueOf(q.ViewRefreshIcon),
			"ViewRestoreIcon":           reflect.ValueOf(q.ViewRestoreIcon),
			"VisibilityIcon":            reflect.ValueOf(q.VisibilityIcon),
			"VisibilityOffIcon":         reflect.ValueOf(q.VisibilityOffIcon),
			"VolumeDownIcon":            reflect.ValueOf(q.VolumeDownIcon),
			"VolumeMuteIcon":            reflect.ValueOf(q.VolumeMuteIcon),
			"VolumeUpIcon":              reflect.ValueOf(q.VolumeUpIcon),
			"WarningColor":              reflect.ValueOf(q.WarningColor),
			"WarningIcon":               reflect.ValueOf(q.WarningIcon),
			"ZoomFitIcon":               reflect.ValueOf(q.ZoomFitIcon),
			"ZoomInIcon":                reflect.ValueOf(q.ZoomInIcon),
			"ZoomOutIcon":               reflect.ValueOf(q.ZoomOutIcon),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ColorNameBackground":        {reflect.TypeOf(q.ColorNameBackground), constant.MakeString(string(q.ColorNameBackground))},
			"ColorNameButton":            {reflect.TypeOf(q.ColorNameButton), constant.MakeString(string(q.ColorNameButton))},
			"ColorNameDisabled":          {reflect.TypeOf(q.ColorNameDisabled), constant.MakeString(string(q.ColorNameDisabled))},
			"ColorNameDisabledButton":    {reflect.TypeOf(q.ColorNameDisabledButton), constant.MakeString(string(q.ColorNameDisabledButton))},
			"ColorNameError":             {reflect.TypeOf(q.ColorNameError), constant.MakeString(string(q.ColorNameError))},
			"ColorNameFocus":             {reflect.TypeOf(q.ColorNameFocus), constant.MakeString(string(q.ColorNameFocus))},
			"ColorNameForeground":        {reflect.TypeOf(q.ColorNameForeground), constant.MakeString(string(q.ColorNameForeground))},
			"ColorNameHover":             {reflect.TypeOf(q.ColorNameHover), constant.MakeString(string(q.ColorNameHover))},
			"ColorNameInputBackground":   {reflect.TypeOf(q.ColorNameInputBackground), constant.MakeString(string(q.ColorNameInputBackground))},
			"ColorNameInputBorder":       {reflect.TypeOf(q.ColorNameInputBorder), constant.MakeString(string(q.ColorNameInputBorder))},
			"ColorNameMenuBackground":    {reflect.TypeOf(q.ColorNameMenuBackground), constant.MakeString(string(q.ColorNameMenuBackground))},
			"ColorNameOverlayBackground": {reflect.TypeOf(q.ColorNameOverlayBackground), constant.MakeString(string(q.ColorNameOverlayBackground))},
			"ColorNamePlaceHolder":       {reflect.TypeOf(q.ColorNamePlaceHolder), constant.MakeString(string(q.ColorNamePlaceHolder))},
			"ColorNamePressed":           {reflect.TypeOf(q.ColorNamePressed), constant.MakeString(string(q.ColorNamePressed))},
			"ColorNamePrimary":           {reflect.TypeOf(q.ColorNamePrimary), constant.MakeString(string(q.ColorNamePrimary))},
			"ColorNameScrollBar":         {reflect.TypeOf(q.ColorNameScrollBar), constant.MakeString(string(q.ColorNameScrollBar))},
			"ColorNameSelection":         {reflect.TypeOf(q.ColorNameSelection), constant.MakeString(string(q.ColorNameSelection))},
			"ColorNameSeparator":         {reflect.TypeOf(q.ColorNameSeparator), constant.MakeString(string(q.ColorNameSeparator))},
			"ColorNameShadow":            {reflect.TypeOf(q.ColorNameShadow), constant.MakeString(string(q.ColorNameShadow))},
			"ColorNameSuccess":           {reflect.TypeOf(q.ColorNameSuccess), constant.MakeString(string(q.ColorNameSuccess))},
			"ColorNameWarning":           {reflect.TypeOf(q.ColorNameWarning), constant.MakeString(string(q.ColorNameWarning))},
			"IconNameAccount":            {reflect.TypeOf(q.IconNameAccount), constant.MakeString(string(q.IconNameAccount))},
			"IconNameArrowDropDown":      {reflect.TypeOf(q.IconNameArrowDropDown), constant.MakeString(string(q.IconNameArrowDropDown))},
			"IconNameArrowDropUp":        {reflect.TypeOf(q.IconNameArrowDropUp), constant.MakeString(string(q.IconNameArrowDropUp))},
			"IconNameCancel":             {reflect.TypeOf(q.IconNameCancel), constant.MakeString(string(q.IconNameCancel))},
			"IconNameCheckButton":        {reflect.TypeOf(q.IconNameCheckButton), constant.MakeString(string(q.IconNameCheckButton))},
			"IconNameCheckButtonChecked": {reflect.TypeOf(q.IconNameCheckButtonChecked), constant.MakeString(string(q.IconNameCheckButtonChecked))},
			"IconNameColorAchromatic":    {reflect.TypeOf(q.IconNameColorAchromatic), constant.MakeString(string(q.IconNameColorAchromatic))},
			"IconNameColorChromatic":     {reflect.TypeOf(q.IconNameColorChromatic), constant.MakeString(string(q.IconNameColorChromatic))},
			"IconNameColorPalette":       {reflect.TypeOf(q.IconNameColorPalette), constant.MakeString(string(q.IconNameColorPalette))},
			"IconNameComputer":           {reflect.TypeOf(q.IconNameComputer), constant.MakeString(string(q.IconNameComputer))},
			"IconNameConfirm":            {reflect.TypeOf(q.IconNameConfirm), constant.MakeString(string(q.IconNameConfirm))},
			"IconNameContentAdd":         {reflect.TypeOf(q.IconNameContentAdd), constant.MakeString(string(q.IconNameContentAdd))},
			"IconNameContentClear":       {reflect.TypeOf(q.IconNameContentClear), constant.MakeString(string(q.IconNameContentClear))},
			"IconNameContentCopy":        {reflect.TypeOf(q.IconNameContentCopy), constant.MakeString(string(q.IconNameContentCopy))},
			"IconNameContentCut":         {reflect.TypeOf(q.IconNameContentCut), constant.MakeString(string(q.IconNameContentCut))},
			"IconNameContentPaste":       {reflect.TypeOf(q.IconNameContentPaste), constant.MakeString(string(q.IconNameContentPaste))},
			"IconNameContentRedo":        {reflect.TypeOf(q.IconNameContentRedo), constant.MakeString(string(q.IconNameContentRedo))},
			"IconNameContentRemove":      {reflect.TypeOf(q.IconNameContentRemove), constant.MakeString(string(q.IconNameContentRemove))},
			"IconNameContentUndo":        {reflect.TypeOf(q.IconNameContentUndo), constant.MakeString(string(q.IconNameContentUndo))},
			"IconNameDelete":             {reflect.TypeOf(q.IconNameDelete), constant.MakeString(string(q.IconNameDelete))},
			"IconNameDocument":           {reflect.TypeOf(q.IconNameDocument), constant.MakeString(string(q.IconNameDocument))},
			"IconNameDocumentCreate":     {reflect.TypeOf(q.IconNameDocumentCreate), constant.MakeString(string(q.IconNameDocumentCreate))},
			"IconNameDocumentPrint":      {reflect.TypeOf(q.IconNameDocumentPrint), constant.MakeString(string(q.IconNameDocumentPrint))},
			"IconNameDocumentSave":       {reflect.TypeOf(q.IconNameDocumentSave), constant.MakeString(string(q.IconNameDocumentSave))},
			"IconNameDownload":           {reflect.TypeOf(q.IconNameDownload), constant.MakeString(string(q.IconNameDownload))},
			"IconNameError":              {reflect.TypeOf(q.IconNameError), constant.MakeString(string(q.IconNameError))},
			"IconNameFile":               {reflect.TypeOf(q.IconNameFile), constant.MakeString(string(q.IconNameFile))},
			"IconNameFileApplication":    {reflect.TypeOf(q.IconNameFileApplication), constant.MakeString(string(q.IconNameFileApplication))},
			"IconNameFileAudio":          {reflect.TypeOf(q.IconNameFileAudio), constant.MakeString(string(q.IconNameFileAudio))},
			"IconNameFileImage":          {reflect.TypeOf(q.IconNameFileImage), constant.MakeString(string(q.IconNameFileImage))},
			"IconNameFileText":           {reflect.TypeOf(q.IconNameFileText), constant.MakeString(string(q.IconNameFileText))},
			"IconNameFileVideo":          {reflect.TypeOf(q.IconNameFileVideo), constant.MakeString(string(q.IconNameFileVideo))},
			"IconNameFolder":             {reflect.TypeOf(q.IconNameFolder), constant.MakeString(string(q.IconNameFolder))},
			"IconNameFolderNew":          {reflect.TypeOf(q.IconNameFolderNew), constant.MakeString(string(q.IconNameFolderNew))},
			"IconNameFolderOpen":         {reflect.TypeOf(q.IconNameFolderOpen), constant.MakeString(string(q.IconNameFolderOpen))},
			"IconNameGrid":               {reflect.TypeOf(q.IconNameGrid), constant.MakeString(string(q.IconNameGrid))},
			"IconNameHelp":               {reflect.TypeOf(q.IconNameHelp), constant.MakeString(string(q.IconNameHelp))},
			"IconNameHistory":            {reflect.TypeOf(q.IconNameHistory), constant.MakeString(string(q.IconNameHistory))},
			"IconNameHome":               {reflect.TypeOf(q.IconNameHome), constant.MakeString(string(q.IconNameHome))},
			"IconNameInfo":               {reflect.TypeOf(q.IconNameInfo), constant.MakeString(string(q.IconNameInfo))},
			"IconNameList":               {reflect.TypeOf(q.IconNameList), constant.MakeString(string(q.IconNameList))},
			"IconNameLogin":              {reflect.TypeOf(q.IconNameLogin), constant.MakeString(string(q.IconNameLogin))},
			"IconNameLogout":             {reflect.TypeOf(q.IconNameLogout), constant.MakeString(string(q.IconNameLogout))},
			"IconNameMailAttachment":     {reflect.TypeOf(q.IconNameMailAttachment), constant.MakeString(string(q.IconNameMailAttachment))},
			"IconNameMailCompose":        {reflect.TypeOf(q.IconNameMailCompose), constant.MakeString(string(q.IconNameMailCompose))},
			"IconNameMailForward":        {reflect.TypeOf(q.IconNameMailForward), constant.MakeString(string(q.IconNameMailForward))},
			"IconNameMailReply":          {reflect.TypeOf(q.IconNameMailReply), constant.MakeString(string(q.IconNameMailReply))},
			"IconNameMailReplyAll":       {reflect.TypeOf(q.IconNameMailReplyAll), constant.MakeString(string(q.IconNameMailReplyAll))},
			"IconNameMailSend":           {reflect.TypeOf(q.IconNameMailSend), constant.MakeString(string(q.IconNameMailSend))},
			"IconNameMediaFastForward":   {reflect.TypeOf(q.IconNameMediaFastForward), constant.MakeString(string(q.IconNameMediaFastForward))},
			"IconNameMediaFastRewind":    {reflect.TypeOf(q.IconNameMediaFastRewind), constant.MakeString(string(q.IconNameMediaFastRewind))},
			"IconNameMediaMusic":         {reflect.TypeOf(q.IconNameMediaMusic), constant.MakeString(string(q.IconNameMediaMusic))},
			"IconNameMediaPause":         {reflect.TypeOf(q.IconNameMediaPause), constant.MakeString(string(q.IconNameMediaPause))},
			"IconNameMediaPhoto":         {reflect.TypeOf(q.IconNameMediaPhoto), constant.MakeString(string(q.IconNameMediaPhoto))},
			"IconNameMediaPlay":          {reflect.TypeOf(q.IconNameMediaPlay), constant.MakeString(string(q.IconNameMediaPlay))},
			"IconNameMediaRecord":        {reflect.TypeOf(q.IconNameMediaRecord), constant.MakeString(string(q.IconNameMediaRecord))},
			"IconNameMediaReplay":        {reflect.TypeOf(q.IconNameMediaReplay), constant.MakeString(string(q.IconNameMediaReplay))},
			"IconNameMediaSkipNext":      {reflect.TypeOf(q.IconNameMediaSkipNext), constant.MakeString(string(q.IconNameMediaSkipNext))},
			"IconNameMediaSkipPrevious":  {reflect.TypeOf(q.IconNameMediaSkipPrevious), constant.MakeString(string(q.IconNameMediaSkipPrevious))},
			"IconNameMediaStop":          {reflect.TypeOf(q.IconNameMediaStop), constant.MakeString(string(q.IconNameMediaStop))},
			"IconNameMediaVideo":         {reflect.TypeOf(q.IconNameMediaVideo), constant.MakeString(string(q.IconNameMediaVideo))},
			"IconNameMenu":               {reflect.TypeOf(q.IconNameMenu), constant.MakeString(string(q.IconNameMenu))},
			"IconNameMenuExpand":         {reflect.TypeOf(q.IconNameMenuExpand), constant.MakeString(string(q.IconNameMenuExpand))},
			"IconNameMoreHorizontal":     {reflect.TypeOf(q.IconNameMoreHorizontal), constant.MakeString(string(q.IconNameMoreHorizontal))},
			"IconNameMoreVertical":       {reflect.TypeOf(q.IconNameMoreVertical), constant.MakeString(string(q.IconNameMoreVertical))},
			"IconNameMoveDown":           {reflect.TypeOf(q.IconNameMoveDown), constant.MakeString(string(q.IconNameMoveDown))},
			"IconNameMoveUp":             {reflect.TypeOf(q.IconNameMoveUp), constant.MakeString(string(q.IconNameMoveUp))},
			"IconNameNavigateBack":       {reflect.TypeOf(q.IconNameNavigateBack), constant.MakeString(string(q.IconNameNavigateBack))},
			"IconNameNavigateNext":       {reflect.TypeOf(q.IconNameNavigateNext), constant.MakeString(string(q.IconNameNavigateNext))},
			"IconNameQuestion":           {reflect.TypeOf(q.IconNameQuestion), constant.MakeString(string(q.IconNameQuestion))},
			"IconNameRadioButton":        {reflect.TypeOf(q.IconNameRadioButton), constant.MakeString(string(q.IconNameRadioButton))},
			"IconNameRadioButtonChecked": {reflect.TypeOf(q.IconNameRadioButtonChecked), constant.MakeString(string(q.IconNameRadioButtonChecked))},
			"IconNameSearch":             {reflect.TypeOf(q.IconNameSearch), constant.MakeString(string(q.IconNameSearch))},
			"IconNameSearchReplace":      {reflect.TypeOf(q.IconNameSearchReplace), constant.MakeString(string(q.IconNameSearchReplace))},
			"IconNameSettings":           {reflect.TypeOf(q.IconNameSettings), constant.MakeString(string(q.IconNameSettings))},
			"IconNameStorage":            {reflect.TypeOf(q.IconNameStorage), constant.MakeString(string(q.IconNameStorage))},
			"IconNameUpload":             {reflect.TypeOf(q.IconNameUpload), constant.MakeString(string(q.IconNameUpload))},
			"IconNameViewFullScreen":     {reflect.TypeOf(q.IconNameViewFullScreen), constant.MakeString(string(q.IconNameViewFullScreen))},
			"IconNameViewRefresh":        {reflect.TypeOf(q.IconNameViewRefresh), constant.MakeString(string(q.IconNameViewRefresh))},
			"IconNameViewRestore":        {reflect.TypeOf(q.IconNameViewRestore), constant.MakeString(string(q.IconNameViewRestore))},
			"IconNameViewZoomFit":        {reflect.TypeOf(q.IconNameViewZoomFit), constant.MakeString(string(q.IconNameViewZoomFit))},
			"IconNameViewZoomIn":         {reflect.TypeOf(q.IconNameViewZoomIn), constant.MakeString(string(q.IconNameViewZoomIn))},
			"IconNameViewZoomOut":        {reflect.TypeOf(q.IconNameViewZoomOut), constant.MakeString(string(q.IconNameViewZoomOut))},
			"IconNameVisibility":         {reflect.TypeOf(q.IconNameVisibility), constant.MakeString(string(q.IconNameVisibility))},
			"IconNameVisibilityOff":      {reflect.TypeOf(q.IconNameVisibilityOff), constant.MakeString(string(q.IconNameVisibilityOff))},
			"IconNameVolumeDown":         {reflect.TypeOf(q.IconNameVolumeDown), constant.MakeString(string(q.IconNameVolumeDown))},
			"IconNameVolumeMute":         {reflect.TypeOf(q.IconNameVolumeMute), constant.MakeString(string(q.IconNameVolumeMute))},
			"IconNameVolumeUp":           {reflect.TypeOf(q.IconNameVolumeUp), constant.MakeString(string(q.IconNameVolumeUp))},
			"IconNameWarning":            {reflect.TypeOf(q.IconNameWarning), constant.MakeString(string(q.IconNameWarning))},
			"SizeNameCaptionText":        {reflect.TypeOf(q.SizeNameCaptionText), constant.MakeString(string(q.SizeNameCaptionText))},
			"SizeNameHeadingText":        {reflect.TypeOf(q.SizeNameHeadingText), constant.MakeString(string(q.SizeNameHeadingText))},
			"SizeNameInlineIcon":         {reflect.TypeOf(q.SizeNameInlineIcon), constant.MakeString(string(q.SizeNameInlineIcon))},
			"SizeNameInnerPadding":       {reflect.TypeOf(q.SizeNameInnerPadding), constant.MakeString(string(q.SizeNameInnerPadding))},
			"SizeNameInputBorder":        {reflect.TypeOf(q.SizeNameInputBorder), constant.MakeString(string(q.SizeNameInputBorder))},
			"SizeNameLineSpacing":        {reflect.TypeOf(q.SizeNameLineSpacing), constant.MakeString(string(q.SizeNameLineSpacing))},
			"SizeNamePadding":            {reflect.TypeOf(q.SizeNamePadding), constant.MakeString(string(q.SizeNamePadding))},
			"SizeNameScrollBar":          {reflect.TypeOf(q.SizeNameScrollBar), constant.MakeString(string(q.SizeNameScrollBar))},
			"SizeNameScrollBarSmall":     {reflect.TypeOf(q.SizeNameScrollBarSmall), constant.MakeString(string(q.SizeNameScrollBarSmall))},
			"SizeNameSeparatorThickness": {reflect.TypeOf(q.SizeNameSeparatorThickness), constant.MakeString(string(q.SizeNameSeparatorThickness))},
			"SizeNameSubHeadingText":     {reflect.TypeOf(q.SizeNameSubHeadingText), constant.MakeString(string(q.SizeNameSubHeadingText))},
			"SizeNameText":               {reflect.TypeOf(q.SizeNameText), constant.MakeString(string(q.SizeNameText))},
			"VariantDark":                {reflect.TypeOf(q.VariantDark), constant.MakeInt64(int64(q.VariantDark))},
			"VariantLight":               {reflect.TypeOf(q.VariantLight), constant.MakeInt64(int64(q.VariantLight))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"ColorBlue":   {"untyped string", constant.MakeString(string(q.ColorBlue))},
			"ColorBrown":  {"untyped string", constant.MakeString(string(q.ColorBrown))},
			"ColorGray":   {"untyped string", constant.MakeString(string(q.ColorGray))},
			"ColorGreen":  {"untyped string", constant.MakeString(string(q.ColorGreen))},
			"ColorOrange": {"untyped string", constant.MakeString(string(q.ColorOrange))},
			"ColorPurple": {"untyped string", constant.MakeString(string(q.ColorPurple))},
			"ColorRed":    {"untyped string", constant.MakeString(string(q.ColorRed))},
			"ColorYellow": {"untyped string", constant.MakeString(string(q.ColorYellow))},
		},
	})
}
