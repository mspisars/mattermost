// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import {renderWithContext, screen} from 'tests/react_testing_utils';
import {RHSStates} from 'utils/constants';

import type {RhsState} from 'types/store/rhs';

import Header from './search_results_header';

describe('search_results_header', () => {
    test('should display back button when the parent is channel info', () => {
        renderWithContext(
            <Header
                previousRhsState={RHSStates.CHANNEL_INFO as RhsState}
                canGoBack={true}
                isExpanded={false}
                actions={{
                    closeRightHandSide: jest.fn(),
                    toggleRhsExpanded: jest.fn(),
                    goBack: jest.fn(),
                }}
            >
                {'Title'}
            </Header>,
        );

        expect(screen.getByLabelText('Go back')).toBeInTheDocument();
    });
    test('should NOT diplay expand when the parent is channel info', () => {
        renderWithContext(
            <Header
                previousRhsState={RHSStates.CHANNEL_INFO as RhsState}
                canGoBack={true}
                isExpanded={false}
                actions={{
                    closeRightHandSide: jest.fn(),
                    toggleRhsExpanded: jest.fn(),
                    goBack: jest.fn(),
                }}
            >
                {'Title'}
            </Header>,
        );

        expect(screen.queryByLabelText('Expand sidebar')).not.toBeInTheDocument();
    });
    test('should diplay expand when the parent is NOT channel info', () => {
        renderWithContext(
            <Header
                previousRhsState={RHSStates.FLAG as RhsState}
                canGoBack={true}
                isExpanded={false}
                actions={{
                    closeRightHandSide: jest.fn(),
                    toggleRhsExpanded: jest.fn(),
                    goBack: jest.fn(),
                }}
            >
                {'Title'}
            </Header>,
        );

        expect(screen.getByLabelText('Expand sidebar')).toBeInTheDocument();
    });
});
