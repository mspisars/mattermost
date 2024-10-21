// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {memo, useEffect, useCallback, useMemo} from 'react';
import {FormattedMessage, useIntl} from 'react-intl';

import Toast from 'components/toast/toast';

import Constants from 'utils/constants';
import {isKeyPressed} from 'utils/keyboard';

import './new_replies_banner.scss';

type Props = {
    hasNewReplies: boolean;
    width: number;
    onClick: () => void;
    onDismiss: () => void;
    actions: {
        updateThreadToastStatus: (status: boolean) => void;
    };
}

function NewRepliesBanner({
    hasNewReplies,
    onClick,
    onDismiss,
    width,
    actions: {
        updateThreadToastStatus,
    },
}: Props) {
    const intl = useIntl();
    const onClickMessage = intl.formatMessage({id: 'postlist.toast.scrollToLatest', defaultMessage: 'Jump to new messages'});

    useEffect(() => {
        updateThreadToastStatus(hasNewReplies);
    }, [hasNewReplies]);

    const handleShortcut = useCallback((e) => {
        if (isKeyPressed(e, Constants.KeyCodes.ESCAPE)) {
            if (hasNewReplies) {
                onDismiss();
            }
        }
    }, [hasNewReplies, onDismiss]);

    useEffect(() => {
        document.addEventListener('keydown', handleShortcut);

        return () => {
            document.removeEventListener('keydown', handleShortcut);
        };
    }, [handleShortcut]);

    const style = useMemo(() => ({width}), [width]);
    return (
        <div
            className='new-replies-banner'
            style={style}
        >
            <Toast
                show={hasNewReplies}
                showActions={true}
                onClick={onClick}
                onDismiss={onDismiss}
                onClickMessage={onClickMessage}
                overlayPlacement='top'
                width={156}
            >
                <FormattedMessage
                    id='rhs_thread.toast.newReplies'
                    defaultMessage='New Replies'
                />
            </Toast>
        </div>
    );
}

export default memo(NewRepliesBanner);
