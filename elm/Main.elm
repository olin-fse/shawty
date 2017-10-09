import Html exposing (Html, Attribute, div, input, button, text)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput, onClick)
import Http
import Json.Decode as Decode
import Json.Encode exposing (..)
import Debug exposing (log)

main =
    Html.program 
        { init = init ""
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

-- MODEL

type alias Model =
    { userUrl : String
    , generatedUrl: String
    }

init : String -> (Model, Cmd Msg)
init userUrl = 
    ( Model userUrl ""
    , Cmd.none
    )

-- UPDATE

type Msg
    = Shorten
    | ShowUrl (Result Http.Error String)
    | Change String

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        Shorten ->
            (model, shortenUrl model.userUrl)

        Change newContent ->
            ({ model | userUrl = newContent }, Cmd.none)

        ShowUrl (Ok newUrl) ->
            ({ model | generatedUrl = newUrl }, Cmd.none)

        ShowUrl (Err _) ->
            (model, Cmd.none)


-- HTTP

shortenUrl : String -> Cmd Msg
shortenUrl url =
    let
        request =
            Http.post "/generate" (Http.jsonBody (encodeUserUrl url)) decodeGeneratedUrl
    in
        Http.send ShowUrl request

encodeUserUrl : String -> Value
encodeUserUrl url =
    object
        [ ("Url", string url) 
        ]
            
decodeGeneratedUrl : Decode.Decoder String
decodeGeneratedUrl =
    Decode.at ["Url"] Decode.string

-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ input [ placeholder "Url to shorten", onInput Change, inputStyle ] []
        , button [ onClick Shorten ] [ text "Shorten Url" ]
        , div [] [ text (model.generatedUrl) ]
        ]

inputStyle =
    style
    [ ("width", "300px")
    ]

-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none




