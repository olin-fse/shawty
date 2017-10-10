module View exposing (..)

import Html exposing (Html, Attribute, div, input, button, text)
import Html.Attributes exposing (..)
import Html.Events exposing (onInput, onClick)

import Msgs exposing (Msg(..))
import Models exposing (Model)

view : Model -> Html Msg
view model =
    div []
        [ input [ placeholder "Url to shorten", onInput Change, inputStyle ] []
        , button [ onClick Shorten ] [ text "Shorten Url" ]
        , div [] [ text (model.generatedUrl) ]
        ]

inputStyle : Attribute msg
inputStyle =
    style
    [ ("width", "300px")
    ]
