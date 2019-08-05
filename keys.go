// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojay

package protocol

const (
	keyActions                           = "actions"
	keyActiveParameter                   = "activeParameter"
	keyActiveSignature                   = "activeSignature"
	keyAdded                             = "added"
	keyAdditionalTextEdits               = "additionalTextEdits"
	keyAlpha                             = "alpha"
	keyApplied                           = "applied"
	keyApplyEdit                         = "applyEdit"
	keyArguments                         = "arguments"
	keyBlue                              = "blue"
	keyCapabilities                      = "capabilities"
	keyCh                                = "ch"
	keyChange                            = "change"
	keyChangeNotifications               = "changeNotifications"
	keyChanges                           = "changes"
	keyCharacter                         = "character"
	keyChildren                          = "children"
	keyCode                              = "code"
	keyCodeAction                        = "codeAction"
	keyCodeActionKind                    = "codeActionKind"
	keyCodeActionKinds                   = "codeActionKinds"
	keyCodeActionLiteralSupport          = "codeActionLiteralSupport"
	keyCodeActionProvider                = "codeActionProvider"
	keyCodeLens                          = "codeLens"
	keyCodeLensProvider                  = "codeLensProvider"
	keyColor                             = "color"
	keyColorProvider                     = "colorProvider"
	keyCommand                           = "command"
	keyCommands                          = "commands"
	keyCommitCharacters                  = "commitCharacters"
	keyCommitCharactersSupport           = "commitCharactersSupport"
	keyCompletion                        = "completion"
	keyCompletionItem                    = "completionItem"
	keyCompletionItemKind                = "completionItemKind"
	keyCompletionProvider                = "completionProvider"
	keyConfiguration                     = "configuration"
	keyContainerName                     = "containerName"
	keyContentChanges                    = "contentChanges"
	keyContentFormat                     = "contentFormat"
	keyContents                          = "contents"
	keyContext                           = "context"
	keyContextSupport                    = "contextSupport"
	keyData                              = "data"
	keyDeclaration                       = "declaration"
	keyDefinition                        = "definition"
	keyDefinitionProvider                = "definitionProvider"
	keyDeprecated                        = "deprecated"
	keyDeprecatedSupport                 = "deprecatedSupport"
	keyDetail                            = "detail"
	keyDiagnostics                       = "diagnostics"
	keyDidChangeConfiguration            = "didChangeConfiguration"
	keyDidChangeWatchedFiles             = "didChangeWatchedFiles"
	keyDidSave                           = "didSave"
	keyDocument                          = "document"
	keyDocumentation                     = "documentation"
	keyDocumentationFormat               = "documentationFormat"
	keyDocumentChanges                   = "documentChanges"
	keyDocumentFormattingProvider        = "documentFormattingProvider"
	keyDocumentHighlight                 = "documentHighlight"
	keyDocumentHighlightProvider         = "documentHighlightProvider"
	keyDocumentLink                      = "documentLink"
	keyDocumentLinkProvider              = "documentLinkProvider"
	keyDocumentOnTypeFormattingProvider  = "documentOnTypeFormattingProvider"
	keyDocumentRangeFormattingProvider   = "documentRangeFormattingProvider"
	keyDocumentSelector                  = "documentSelector"
	keyDocumentSymbol                    = "documentSymbol"
	keyDocumentSymbolProvider            = "documentSymbolProvider"
	keyDynamicRegistration               = "dynamicRegistration"
	keyEdit                              = "edit"
	keyEdits                             = "edits"
	keyEnd                               = "end"
	keyEndCharacter                      = "endCharacter"
	keyEndLine                           = "endLine"
	keyEvent                             = "event"
	keyExecuteCommand                    = "executeCommand"
	keyExecuteCommandProvider            = "executeCommandProvider"
	keyExperimental                      = "experimental"
	keyFailureHandling                   = "failureHandling"
	keyFilterText                        = "filterText"
	keyFirstTriggerCharacter             = "firstTriggerCharacter"
	keyFoldingRange                      = "foldingRange"
	keyFoldingRangeProvider              = "foldingRangeProvider"
	keyFormatting                        = "formatting"
	keyGlobPattern                       = "globPattern"
	keyGreen                             = "green"
	keyHierarchicalDocumentSymbolSupport = "hierarchicalDocumentSymbolSupport"
	keyHover                             = "hover"
	keyHoverProvider                     = "hoverProvider"
	keyID                                = "id"
	keyIgnoreIfExists                    = "ignoreIfExists"
	keyIgnoreIfNotExists                 = "ignoreIfNotExists"
	keyImplementation                    = "implementation"
	keyImplementationProvider            = "implementationProvider"
	keyIncludeDeclaration                = "includeDeclaration"
	keyIncludeText                       = "includeText"
	keyInitializationOptions             = "initializationOptions"
	keyInsertSpaces                      = "insertSpaces"
	keyInsertText                        = "insertText"
	keyInsertTextFormat                  = "insertTextFormat"
	keyIsIncomplete                      = "isIncomplete"
	keyItems                             = "items"
	keyKind                              = "kind"
	keyLabel                             = "label"
	keyLanguage                          = "language"
	keyLanguageID                        = "languageId"
	keyLine                              = "line"
	keyLineCount                         = "lineCount"
	keyLineFoldingOnly                   = "lineFoldingOnly"
	keyLinkSupport                       = "linkSupport"
	keyLocation                          = "location"
	keyMessage                           = "message"
	keyMethod                            = "method"
	keyMoreTriggerCharacter              = "moreTriggerCharacter"
	keyName                              = "name"
	keyNewName                           = "newName"
	keyNewText                           = "newText"
	keyNewURI                            = "newUri"
	keyOldURI                            = "oldUri"
	keyOnly                              = "only"
	keyOnTypeFormatting                  = "onTypeFormatting"
	keyOpenClose                         = "openClose"
	keyOptions                           = "options"
	keyOriginSelectionRange              = "originSelectionRange"
	keyOverwrite                         = "overwrite"
	keyParameterInformation              = "parameterInformation"
	keyPattern                           = "pattern"
	keyPosition                          = "position"
	keyPrepareProvider                   = "prepareProvider"
	keyPrepareSupport                    = "prepareSupport"
	keyPreselect                         = "preselect"
	keyPreselectSupport                  = "preselectSupport"
	keyProcessID                         = "processId"
	keyPublishDiagnostics                = "publishDiagnostics"
	keyQuery                             = "query"
	keyRange                             = "range"
	keyRangeFormatting                   = "rangeFormatting"
	keyRangeLength                       = "rangeLength"
	keyRangeLimit                        = "rangeLimit"
	keyReason                            = "reason"
	keyRecursive                         = "recursive"
	keyRed                               = "red"
	keyReferences                        = "references"
	keyReferencesProvider                = "referencesProvider"
	keyRegisterOptions                   = "registerOptions"
	keyRegistrations                     = "registrations"
	keyRelatedInformation                = "relatedInformation"
	keyRemoved                           = "removed"
	keyRename                            = "rename"
	keyRenameProvider                    = "renameProvider"
	keyResolveProvider                   = "resolveProvider"
	keyResourceOperations                = "resourceOperations"
	keyRetry                             = "retry"
	keyRootPath                          = "rootPath"
	keyRootURI                           = "rootUri"
	keySave                              = "save"
	keyScheme                            = "scheme"
	keyScopeURI                          = "scopeUri"
	keySection                           = "section"
	keySelectionRange                    = "selectionRange"
	keySelectionRangeProvider            = "selectionRangeProvider"
	keySettings                          = "settings"
	keySeverity                          = "severity"
	keySignatureHelp                     = "signatureHelp"
	keySignatureHelpProvider             = "signatureHelpProvider"
	keySignatureInformation              = "signatureInformation"
	keySignatures                        = "signatures"
	keySnippetSupport                    = "snippetSupport"
	keySortText                          = "sortText"
	keySource                            = "source"
	keyStart                             = "start"
	keyStartCharacter                    = "startCharacter"
	keyStartLine                         = "startLine"
	keySupported                         = "supported"
	keySymbol                            = "symbol"
	keySymbolKind                        = "symbolKind"
	keySynchronization                   = "synchronization"
	keySyncKind                          = "syncKind"
	keyTabSize                           = "tabSize"
	keyTagSupport                        = "keyTagSupport"
	keyTarget                            = "target"
	keyTargetRange                       = "targetRange"
	keyTargetSelectionRange              = "targetSelectionRange"
	keyTargetURI                         = "targetUri"
	keyText                              = "text"
	keyTextDocument                      = "textDocument"
	keyTextDocumentSync                  = "textDocumentSync"
	keyTextEdit                          = "textEdit"
	keyTitle                             = "title"
	keyTrace                             = "trace"
	keyTriggerCharacter                  = "triggerCharacter"
	keyTriggerCharacters                 = "triggerCharacters"
	keyTriggerKind                       = "triggerKind"
	keyType                              = "type"
	keyTypeDefinition                    = "typeDefinition"
	keyTypeDefinitionProvider            = "typeDefinitionProvider"
	keyUnregisterations                  = "unregisterations"
	keyURI                               = "uri"
	keyValue                             = "value"
	keyValueSet                          = "valueSet"
	keyVersion                           = "version"
	keyWatchers                          = "watchers"
	keyWillSave                          = "willSave"
	keyWillSaveWaitUntil                 = "willSaveWaitUntil"
	keyWorkspace                         = "workspace"
	keyWorkspaceEdit                     = "workspaceEdit"
	keyWorkspaceFolders                  = "workspaceFolders"
	keyWorkspaceSymbolProvider           = "workspaceSymbolProvider"
)
