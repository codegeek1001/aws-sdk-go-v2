// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Sends user input to Amazon Lex V2. You can send text or speech. Clients use this
// API to send text and audio requests to Amazon Lex V2 at runtime. Amazon Lex V2
// interprets the user input using the machine learning model built for the bot.
// The following request fields must be compressed with gzip and then base64
// encoded before you send them to Amazon Lex V2.
//
// * requestAttributes
//
// *
// sessionState
//
// The following response fields are compressed using gzip and then
// base64 encoded by Amazon Lex V2. Before you can use these fields, you must
// decode and decompress them.
//
// * inputTranscript
//
// * interpretations
//
// * messages
//
// *
// requestAttributes
//
// * sessionState
//
// The example contains a Java application that
// compresses and encodes a Java object to send to Amazon Lex V2, and a second that
// decodes and decompresses a response from Amazon Lex V2.
func (c *Client) RecognizeUtterance(ctx context.Context, params *RecognizeUtteranceInput, optFns ...func(*Options)) (*RecognizeUtteranceOutput, error) {
	if params == nil {
		params = &RecognizeUtteranceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecognizeUtterance", params, optFns, c.addOperationRecognizeUtteranceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecognizeUtteranceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecognizeUtteranceInput struct {

	// The alias identifier in use for the bot that should receive the request.
	//
	// This member is required.
	BotAliasId *string

	// The identifier of the bot that should receive the request.
	//
	// This member is required.
	BotId *string

	// The locale where the session is in use.
	//
	// This member is required.
	LocaleId *string

	// Indicates the format for audio input or that the content is text. The header
	// must start with one of the following prefixes:
	//
	// * PCM format, audio data must be
	// in little-endian byte order.
	//
	// * audio/l16; rate=16000; channels=1
	//
	// *
	// audio/x-l16; sample-rate=16000; channel-count=1
	//
	// * audio/lpcm; sample-rate=8000;
	// sample-size-bits=16; channel-count=1; is-big-endian=false
	//
	// * Opus format
	//
	// *
	// audio/x-cbr-opus-with-preamble;preamble-size=0;bit-rate=256000;frame-size-milliseconds=4
	//
	// *
	// Text format
	//
	// * text/plain; charset=utf-8
	//
	// This member is required.
	RequestContentType *string

	// The identifier of the session in use.
	//
	// This member is required.
	SessionId *string

	// User input in PCM or Opus audio format or text format as described in the
	// requestContentType parameter.
	InputStream io.Reader

	// Request-specific information passed between the client application and Amazon
	// Lex V2 The namespace x-amz-lex: is reserved for special attributes. Don't create
	// any request attributes for prefix x-amz-lex:. The requestAttributes field must
	// be compressed using gzip and then base64 encoded before sending to Amazon Lex
	// V2.
	RequestAttributes *string

	// The message that Amazon Lex V2 returns in the response can be either text or
	// speech based on the responseContentType value.
	//
	// * If the value is
	// text/plain;charset=utf-8, Amazon Lex V2 returns text in the response.
	//
	// * If the
	// value begins with audio/, Amazon Lex V2 returns speech in the response. Amazon
	// Lex V2 uses Amazon Polly to generate the speech using the configuration that you
	// specified in the requestContentType parameter. For example, if you specify
	// audio/mpeg as the value, Amazon Lex V2 returns speech in the MPEG format.
	//
	// * If
	// the value is audio/pcm, the speech returned is audio/pcm at 16 KHz in 16-bit,
	// little-endian format.
	//
	// * The following are the accepted values:
	//
	// * audio/mpeg
	//
	// *
	// audio/ogg
	//
	// * audio/pcm (16 KHz)
	//
	// * audio/* (defaults to mpeg)
	//
	// * text/plain;
	// charset=utf-8
	ResponseContentType *string

	// Sets the state of the session with the user. You can use this to set the current
	// intent, attributes, context, and dialog action. Use the dialog action to
	// determine the next step that Amazon Lex V2 should use in the conversation with
	// the user. The sessionState field must be compressed using gzip and then base64
	// encoded before sending to Amazon Lex V2.
	SessionState *string

	noSmithyDocumentSerde
}

type RecognizeUtteranceOutput struct {

	// The prompt or statement to send to the user. This is based on the bot
	// configuration and context. For example, if Amazon Lex V2 did not understand the
	// user intent, it sends the clarificationPrompt configured for the bot. If the
	// intent requires confirmation before taking the fulfillment action, it sends the
	// confirmationPrompt. Another example: Suppose that the Lambda function
	// successfully fulfilled the intent, and sent a message to convey to the user.
	// Then Amazon Lex V2 sends that message in the response.
	AudioStream io.ReadCloser

	// Content type as specified in the responseContentType in the request.
	ContentType *string

	// Indicates whether the input mode to the operation was text or speech.
	InputMode *string

	// The text used to process the request. If the input was an audio stream, the
	// inputTranscript field contains the text extracted from the audio stream. This is
	// the text that is actually processed to recognize intents and slot values. You
	// can use this information to determine if Amazon Lex V2 is correctly processing
	// the audio that you send. The inputTranscript field is compressed with gzip and
	// then base64 encoded. Before you can use the contents of the field, you must
	// decode and decompress the contents. See the example for a simple function to
	// decode and decompress the contents.
	InputTranscript *string

	// A list of intents that Amazon Lex V2 determined might satisfy the user's
	// utterance. Each interpretation includes the intent, a score that indicates how
	// confident Amazon Lex V2 is that the interpretation is the correct one, and an
	// optional sentiment response that indicates the sentiment expressed in the
	// utterance. The interpretations field is compressed with gzip and then base64
	// encoded. Before you can use the contents of the field, you must decode and
	// decompress the contents. See the example for a simple function to decode and
	// decompress the contents.
	Interpretations *string

	// A list of messages that were last sent to the user. The messages are ordered
	// based on the order that you returned the messages from your Lambda function or
	// the order that the messages are defined in the bot. The messages field is
	// compressed with gzip and then base64 encoded. Before you can use the contents of
	// the field, you must decode and decompress the contents. See the example for a
	// simple function to decode and decompress the contents.
	Messages *string

	// The attributes sent in the request. The requestAttributes field is compressed
	// with gzip and then base64 encoded. Before you can use the contents of the field,
	// you must decode and decompress the contents.
	RequestAttributes *string

	// The identifier of the session in use.
	SessionId *string

	// Represents the current state of the dialog between the user and the bot. Use
	// this to determine the progress of the conversation and what the next action
	// might be. The sessionState field is compressed with gzip and then base64
	// encoded. Before you can use the contents of the field, you must decode and
	// decompress the contents. See the example for a simple function to decode and
	// decompress the contents.
	SessionState *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRecognizeUtteranceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRecognizeUtterance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRecognizeUtterance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRecognizeUtteranceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRecognizeUtterance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRecognizeUtterance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "RecognizeUtterance",
	}
}
