const path = require('path');
const webpack = require('webpack');
const CompressionPlugin = require('compression-webpack-plugin');

module.exports = {
  devtool: 'eval',
  entry: [
    path.join(__dirname, 'src/js/index.js')
  ],
  output: {
    path: path.resolve(__dirname, 'public/'),
    publicPath: '',
    filename: 'bundle.js'
  },
  plugins: [
    new webpack.optimize.ModuleConcatenationPlugin(),
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: JSON.stringify('production')
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
      sourceMap: true,
      mangle: true,

      compress: {
        warnings: false,
        screw_ie8: true
      },

      output: {
        comments: false,
        beautify: false
      },

      exclude: [/\.min\.js$/gi]
    }),
    new CompressionPlugin({
      asset: '[path].gz[query]',
      algorithm: 'gzip',
      test: /\.js$|\.css$|\.html$/,
      threshold: 10240,
      minRatio: 0
    })
  ],
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        include: path.join(__dirname, 'src/js/'),
        loader: 'babel-loader',
        query: {
          presets: ['react', 'env', 'stage-2']
        }
      },
      {
        test: /\.css/,
        use: [
          {loader: 'style-loader'},
          {loader: 'css-loader', options: {importLoaders: 1}}
        ]
      }
    ]
  }
};
