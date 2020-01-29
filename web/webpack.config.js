const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const TerserJSPlugin = require('terser-webpack-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    optimization: {
        minimizer: [
            new TerserJSPlugin({}),
            new OptimizeCSSAssetsPlugin({
                cssProcessor: require('cssnano'),
                cssProcessorPluginOptions: {
                    preset: ['default', { discardComments: { removeAll: true } }],
                },
                canPrint: true
            })
        ],
    },
    mode: process.env.NODE_ENV === 'production' ? 'production' : 'development',
    entry: ['./src/index.js', './src/styles.css'],
    output: {
        path: path.resolve(__dirname, 'build'),
        filename: 'bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: [
                    { loader: MiniCssExtractPlugin.loader, },
                    { loader: 'css-loader', },
                    { loader: 'postcss-loader', },
                ]
            }
        ]
    },
    devtool: process.env.NODE_ENV === 'production' ? '' : 'inline-source-map',
    devServer: {
        contentBase: './build',
    },
    plugins: [
        new CleanWebpackPlugin(),
        new MiniCssExtractPlugin(),
        new HtmlWebpackPlugin({
            template: 'index.template.html',
            hash: true,
        }),
        new CopyWebpackPlugin([
            { from: 'generator.wasm', to: 'generator.wasm' },
            { from: 'wasm_exec.js', to: 'wasm_exec.js' },
        ]),
    ],
};
